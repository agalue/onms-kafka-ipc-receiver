// @author Alejandro Galue <agalue@opennms.org>

// Package client implements a kafka consumer that works with single or multi-part messages for OpenNMS Sink API messages.
package client

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/netflow"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/rpc"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/sink"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/telemetry"
	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// AvailableParsers list of available parsers
var AvailableParsers = &EnumValue{
	Enum: []string{"heartbeat", "snmp", "syslog", "netflow", "sflow"},
}

// ProcessSinkMessage defines the action to execute after successfully received a Sink message.
// It receives the payload as an array of bytes (usually in XML or JSON format)
type ProcessSinkMessage func(key string, msg []byte)

// Propertites represents an array of string flags
type Propertites []string

func (p *Propertites) String() string {
	return strings.Join(*p, ", ")
}

// Set stores a string flag in the array
func (p *Propertites) Set(value string) error {
	*p = append(*p, value)
	return nil
}

// ipcMessage internal structure that represents an IPC message
type ipcMessage struct {
	chunk   int32
	total   int32
	id      string
	content []byte
}

// KafkaClient defines a simple Kafka consumer client.
type KafkaClient struct {
	Bootstrap string // The Kafka Server Bootstrap string.
	Topic     string // The name of the Kafka Topic.
	GroupID   string // The name of the Consumer Group ID.
	IPC       string // options: rpc, sink.
	Parser    string // options: syslog, snmp, netflow, sflow

	subscriber   *kafka.Subscriber
	msgChannel   <-chan *message.Message
	msgBuffer    map[string][]byte
	chunkTracker map[string]int32
	mutex        *sync.RWMutex
	stopping     bool

	msgProcessed   prometheus.Counter
	chunkProcessed prometheus.Counter
}

// Creates the Kafka Configuration Map.
func (cli *KafkaClient) createConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V2_7_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Session.Timeout = 6 * time.Second
	return config
}

// Initializes all internal variables.
func (cli *KafkaClient) createVariables() {
	cli.msgBuffer = make(map[string][]byte)
	cli.chunkTracker = make(map[string]int32)
	cli.mutex = &sync.RWMutex{}
}

func (cli *KafkaClient) registerCounters() {
	cli.msgProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "onms_ipc_processed_messages_total",
		Help: "The total number of processed messages",
	})
	cli.chunkProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "onms_ipc_processed_chunk_total",
		Help: "The total number of processed chunks",
	})
}

func (cli *KafkaClient) getIpcMessage(msg *message.Message) (*ipcMessage, error) {
	if cli.IPC == "rpc" {
		rpcMsg := &rpc.RpcMessageProto{}
		if err := proto.Unmarshal(msg.Payload, rpcMsg); err != nil {
			return nil, fmt.Errorf("[warn] invalid rpc message received: %v", err)
		}
		cli.chunkProcessed.Inc()
		return &ipcMessage{
			chunk:   rpcMsg.CurrentChunkNumber + 1, // Chunks starts at 0
			total:   rpcMsg.TotalChunks,
			id:      rpcMsg.RpcId,
			content: rpcMsg.RpcContent,
		}, nil
	}
	sinkMsg := &sink.SinkMessage{}
	if err := proto.Unmarshal(msg.Payload, sinkMsg); err != nil {
		return nil, fmt.Errorf("[warn] invalid sink message received: %v", err)
	}
	return &ipcMessage{
		chunk:   sinkMsg.GetCurrentChunkNumber() + 1, // Chunks starts at 0
		total:   sinkMsg.GetTotalChunks(),
		id:      sinkMsg.GetMessageId(),
		content: sinkMsg.GetContent(),
	}, nil
}

// Processes a Kafka message. It return a non-empty slice when the message is complete, otherwise returns nil.
// This is a concurrent safe method.
func (cli *KafkaClient) processMessage(msg *message.Message) []byte {
	// Process IPC Messages
	cli.chunkProcessed.Inc()
	ipcmsg, err := cli.getIpcMessage(msg)
	if err != nil {
		log.Printf("[error] invalid IPC message: %v", err)
		return nil
	}
	if ipcmsg.chunk != ipcmsg.total {
		cli.mutex.Lock()
		if cli.chunkTracker[ipcmsg.id] < ipcmsg.chunk {
			// Adds partial message to the buffer
			cli.msgBuffer[ipcmsg.id] = append(cli.msgBuffer[ipcmsg.id], ipcmsg.content...)
			cli.chunkTracker[ipcmsg.id] = ipcmsg.chunk
		} else {
			log.Printf("[warn] chunk %d from %s was already processed, ignoring...", ipcmsg.chunk, ipcmsg.id)
		}
		cli.mutex.Unlock()
		return nil
	}
	// Retrieve the complete message from the buffer
	var data []byte
	if ipcmsg.total == 1 { // Handle special case chunk == total == 1
		data = ipcmsg.content
	} else {
		cli.mutex.RLock()
		data = append(cli.msgBuffer[ipcmsg.id], ipcmsg.content...)
		cli.mutex.RUnlock()
	}
	cli.bufferCleanup(ipcmsg.id)
	cli.msgProcessed.Inc()
	return data
}

func (cli *KafkaClient) isTelemetry() bool {
	return cli.isNetflow() || cli.isSflow()
}

func (cli *KafkaClient) isSflow() bool {
	return strings.ToLower(cli.Parser) == "sflow"
}

func (cli *KafkaClient) isNetflow() bool {
	return strings.ToLower(cli.Parser) == "netflow"
}

func (cli *KafkaClient) isSyslog() bool {
	return strings.ToLower(cli.Parser) == "syslog"
}

func (cli *KafkaClient) isSnmp() bool {
	return strings.ToLower(cli.Parser) == "snmp"
}

func (cli *KafkaClient) isHeartbeat() bool {
	return strings.ToLower(cli.Parser) == "heartbeat"
}

func (cli *KafkaClient) processPayload(key string, data []byte, action ProcessSinkMessage) {
	if cli.IPC == "rpc" {
		action(key, data)
		return
	}
	// log.Printf("[debug] received %s", string(data))
	if cli.isTelemetry() {
		msgLog := &telemetry.TelemetryMessageLog{}
		if err := proto.Unmarshal(data, msgLog); err != nil {
			log.Printf("[warn] error processing telemetry message: %v", err)
			return
		}
		for _, msg := range msgLog.Message {
			if cli.isNetflow() {
				flow := &netflow.FlowMessage{}
				if err := proto.Unmarshal(msg.Bytes, flow); err != nil {
					log.Printf("[warn] invalid netflow message received: %v", err)
					return
				}
				bytes, _ := json.MarshalIndent(flow, "", "  ")
				action(key, bytes)
			} else if cli.isSflow() {
				log.Println("[warn] sflow has not been implemented")
			} else {
				log.Println("[warn] cannot parse telemetry message due to invalid parser")
			}
		}
	} else if cli.isSyslog() {
		syslog := &SyslogMessageLogDTO{}
		if err := xml.Unmarshal(data, syslog); err != nil {
			log.Printf("[warn] invalid syslog message received: %v", err)
			return
		}
		action(key, []byte(syslog.String()))
	} else if cli.isSnmp() {
		trap := &TrapLogDTO{}
		if err := xml.Unmarshal(data, trap); err != nil {
			log.Printf("[warn] invalid snmp trap message received: %v", err)
			return
		}
		action(key, []byte(trap.String()))
	} else if cli.isHeartbeat() {
		action(key, data)
	} else {
		log.Printf("[error] invalid parser %s, ignoring payload", cli.Parser)
	}
}

// Cleans up the chunk buffer. Should be called after successfully processed all chunks.
// This is a concurrent safe method.
func (cli *KafkaClient) bufferCleanup(id string) {
	cli.mutex.Lock()
	delete(cli.msgBuffer, id)
	delete(cli.chunkTracker, id)
	cli.mutex.Unlock()
}

// Initialize builds the Kafka consumer object and the cache for chunk handling.
func (cli *KafkaClient) Initialize(ctx context.Context) error {
	if cli.msgChannel != nil {
		return fmt.Errorf("consumer already initialized")
	}
	if cli.IPC == "" {
		cli.IPC = "sink"
	} else {
		if cli.IPC != "sink" && cli.IPC != "rpc" {
			return fmt.Errorf("invalid IPC %s. Expected sink, rpc", cli.IPC)
		}
	}
	if cli.Parser != "" {
		if err := AvailableParsers.Set(cli.Parser); err != nil {
			return fmt.Errorf("invalid Sink parser %s. Expected %s", cli.Parser, AvailableParsers.EnumAsString())
		}
	}
	var err error
	log.Printf("[info] creating consumer for topic %s at %s", cli.Topic, cli.Bootstrap)

	cli.subscriber, err = kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{cli.Bootstrap},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: cli.createConfig(),
			ConsumerGroup:         cli.GroupID,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return fmt.Errorf("cannot create consumer: %v", err)
	}
	cli.msgChannel, err = cli.subscriber.Subscribe(ctx, cli.Topic)
	if err != nil {
		return fmt.Errorf("cannot subscribe to topic %s: %v", cli.Topic, err)
	}

	cli.createVariables()
	cli.registerCounters()
	return nil
}

func (cli *KafkaClient) byteCount(b float64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%f B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", b/float64(div), "KMGTPE"[exp])
}

// Start registers the consumer for the chosen topic, and reads messages from it on an infinite loop.
// It is recommended to use it within a Go Routine as it is a blocking operation.
func (cli *KafkaClient) Start(action ProcessSinkMessage) {
	if cli.msgChannel == nil {
		log.Fatal("Consumer not initialized")
	}

	jsonBytes, _ := json.MarshalIndent(cli, "", "  ")
	log.Printf("[info] starting kafka consumer: %s", string(jsonBytes))

	cli.stopping = false
	for msg := range cli.msgChannel {
		if data := cli.processMessage(msg); data != nil {
			cli.processPayload(msg.UUID, data, action)
		}
	}
}
