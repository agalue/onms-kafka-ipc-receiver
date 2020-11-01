// @author Alejandro Galue <agalue@opennms.org>

// Package client implements a kafka consumer that works with single or multi-part messages for OpenNMS Sink API messages.
package client

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/flowdocument"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/netflow"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/rpc"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/sink"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/telemetry"
	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// KafkaConsumer creates an generic interface with the relevant methods from kafka.Consumer
type KafkaConsumer interface {
	Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error
	Poll(timeoutMs int) (event kafka.Event)
	CommitMessage(m *kafka.Message) ([]kafka.TopicPartition, error)
	Close() (err error)
}

// ProcessSinkMessage defines the action to execute after successfully received a Sink message.
// It receives the payload as an array of bytes (usually in XML or JSON format)
type ProcessSinkMessage func(key, msg []byte)

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
	Bootstrap  string      // The Kafka Server Bootstrap string.
	Topic      string      // The name of the Kafka Topic.
	GroupID    string      // The name of the Consumer Group ID.
	Parameters Propertites // List of Kafka Consumer Parameters.
	IPC        string      // options: rpc, sink.
	Parser     string      // options: syslog, snmp, netflow, flowdocument

	consumer     KafkaConsumer
	msgBuffer    map[string][]byte
	chunkTracker map[string]int32
	mutex        *sync.RWMutex
	stopping     bool

	msgProcessed   prometheus.Counter
	chunkProcessed prometheus.Counter
}

// Creates the Kafka Configuration Map.
func (cli *KafkaClient) createConfig() *kafka.ConfigMap {
	config := &kafka.ConfigMap{
		"bootstrap.servers":     cli.Bootstrap,
		"group.id":              cli.GroupID,
		"session.timeout.ms":    6000,
		"broker.address.family": "v4",
	}
	if cli.Parameters != nil {
		for _, kv := range cli.Parameters {
			array := strings.Split(kv, "=")
			if len(array) == 2 {
				if err := config.SetKey(array[0], array[1]); err != nil {
					log.Printf("[error] cannot add consumer config %s: %v", kv, err)
				}
			} else {
				log.Printf("[error] invalid key-value pair %s", kv)
			}
		}
	}
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

func (cli *KafkaClient) getIpcMessage(msg *kafka.Message) (*ipcMessage, error) {
	if cli.IPC == "rpc" {
		rpcMsg := &rpc.RpcMessageProto{}
		if err := proto.Unmarshal(msg.Value, rpcMsg); err != nil {
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
	if err := proto.Unmarshal(msg.Value, sinkMsg); err != nil {
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
func (cli *KafkaClient) processMessage(msg *kafka.Message) []byte {
	// Process non-IPC Messages
	if cli.isFlowDocument() {
		flow := &flowdocument.FlowDocument{}
		if err := proto.Unmarshal(msg.Value, flow); err != nil {
			log.Printf("[warn] invalid netflow enriched document message received: %v", err)
			return nil
		}
		bytes, _ := json.MarshalIndent(flow, "", "  ")
		return bytes
	}
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

func (cli *KafkaClient) isFlowDocument() bool {
	return strings.ToLower(cli.Parser) == "flowdocument"
}

func (cli *KafkaClient) processPayload(key, data []byte, action ProcessSinkMessage) {
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
	} else {
		action(key, data)
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
func (cli *KafkaClient) Initialize() error {
	if cli.consumer != nil {
		return fmt.Errorf("consumer already initialized")
	}
	if cli.IPC == "" {
		cli.IPC = "sink"
	} else {
		if cli.IPC != "sink" && cli.IPC != "rpc" {
			return fmt.Errorf("invalid IPC %s. Expected 'sink' or 'rpc'", cli.IPC)
		}
	}
	var err error
	log.Printf("[info] creating consumer for topic %s at %s", cli.Topic, cli.Bootstrap)
	cli.consumer, err = kafka.NewConsumer(cli.createConfig())
	if err != nil {
		return fmt.Errorf("cannot create consumer: %v", err)
	}
	err = cli.consumer.Subscribe(cli.Topic, nil)
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

func (cli *KafkaClient) showStats(sts *kafka.Stats) {
	// https://github.com/edenhill/librdkafka/blob/master/STATISTICS.md
	var stats map[string]interface{}
	json.Unmarshal([]byte(sts.String()), &stats)
	log.Printf("[stats] %v messages (%v) consumed", stats["rxmsgs"], cli.byteCount(stats["rxmsg_bytes"].(float64)))
}

// Start registers the consumer for the chosen topic, and reads messages from it on an infinite loop.
// It is recommended to use it within a Go Routine as it is a blocking operation.
func (cli *KafkaClient) Start(action ProcessSinkMessage) {
	if cli.consumer == nil {
		log.Fatal("Consumer not initialized")
	}

	jsonBytes, _ := json.MarshalIndent(cli, "", "  ")
	log.Printf("[info] starting kafka consumer: %s", string(jsonBytes))

	cli.stopping = false
	for {
		if cli.stopping {
			return
		}
		event := cli.consumer.Poll(500)
		switch e := event.(type) {
		case *kafka.Message:
			log.Printf("[info] received message of %d bytes at %v", len(e.Value), e.TopicPartition)
			if data := cli.processMessage(e); data != nil {
				cli.processPayload(e.Key, data, action)
			}
			_, err := cli.consumer.CommitMessage(e) // If there are errors on the action, the message won't be reprocessed.
			if err != nil {
				log.Printf("[warn] error committing message: %v", err)
			}
		case kafka.Error:
			log.Printf("[warn] consumer error %v", e)
		case *kafka.Stats:
			cli.showStats(e)
		}
	}
}

// Stop terminates the Kafka consumer and waits for the execution of all pending action handlers.
func (cli *KafkaClient) Stop() {
	log.Println("stopping consumer")
	cli.stopping = true
	cli.consumer.Close()
	log.Println("good bye!")
}
