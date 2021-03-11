// @author Alejandro Galue <agalue@opennms.org>

package client

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/netflow"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/sink"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/telemetry"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"gotest.tools/assert"
)

type mockConsumer struct {
	msgChannel chan *kafka.Message
}

func (mock *mockConsumer) Send(msg *kafka.Message) {
	mock.msgChannel <- msg
}

func (mock *mockConsumer) Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error {
	return nil
}

func (mock *mockConsumer) Poll(timeoutMs int) (event kafka.Event) {
	return <-mock.msgChannel
}

func (mock *mockConsumer) Close() (err error) {
	return nil
}

func TestCreateConfig(t *testing.T) {
	cli := &KafkaClient{
		Bootstrap: "kafka1:9092",
		GroupID:   "GoTest",
		Parameters: Propertites{
			"acks=1",
			"fetch.min.bytes=2048",
			"ssl.truststore.location",
		},
	}
	config := cli.createConfig()

	value, err := config.Get("bootstrap.servers", "unknown")
	assert.NilError(t, err)
	assert.Equal(t, cli.Bootstrap, value.(string))

	value, err = config.Get("group.id", "unknown")
	assert.NilError(t, err)
	assert.Equal(t, cli.GroupID, value.(string))

	value, err = config.Get("acks", "0")
	assert.NilError(t, err)
	assert.Equal(t, "1", value.(string))

	value, err = config.Get("fetch.min.bytes", "0")
	assert.NilError(t, err)
	assert.Equal(t, "2048", value.(string))

	value, err = config.Get("ssl.truststore.location", "unknown")
	assert.NilError(t, err)
	assert.Equal(t, "unknown", value.(string))
}

func TestProcessSingleMessage(t *testing.T) {
	cli := createKafkaClient(nil)
	data := cli.processMessage(buildMessage("0001", 0, 1, []byte("ABC")))
	assert.Equal(t, "ABC", string(data))
}

func TestProcessMultipleMessages(t *testing.T) {
	cli := createKafkaClient(nil)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go runProcessMessageTest(t, wg, cli, fmt.Sprintf("ID%d", i))
	}
	wg.Wait()
}

func TestSyslogParser(t *testing.T) {
	mock := &mockConsumer{
		msgChannel: make(chan *kafka.Message, 1),
	}
	cli := createKafkaClient(mock)
	cli.Parser = "syslog"
	var message string
	go func() {
		cli.Start(func(key, msg []byte) {
			fmt.Printf("Key %s, Value: %s\n", string(key), string(msg))
			message = string(msg)
		})
	}()

	dto := SyslogMessageLogDTO{
		Location:      "Test",
		SourceAddress: "10.0.0.1",
		SourcePort:    514,
		Messages: []SyslogMessageDTO{
			{
				Timestamp: time.Now().String(),
				Content:   []byte(base64.StdEncoding.EncodeToString([]byte("This is a test"))),
			},
		},
	}
	data, err := xml.Marshal(dto)
	assert.NilError(t, err)

	mock.Send(buildMessage("001", 0, 1, data))
	time.Sleep(1 * time.Second)
	assert.Assert(t, strings.Contains(message, "This is a test"))
	cli.Stop()
}

func TestSnmpParser(t *testing.T) {
	mock := &mockConsumer{
		msgChannel: make(chan *kafka.Message, 1),
	}
	cli := createKafkaClient(mock)
	cli.Parser = "snmp"
	var message string
	go func() {
		cli.Start(func(key, msg []byte) {
			fmt.Printf("Key %s, Value: %s\n", string(key), string(msg))
			message = string(msg)
		})
	}()

	dto := TrapLogDTO{
		Location:    "Test",
		TrapAddress: "10.0.0.1",
		SystemID:    "mock01",
		Messages: []TrapDTO{
			{
				AgentAddress: "172.16.0.1",
				Version:      "v2",
				Community:    "public",
				CreationTime: time.Now().Unix(),
				Timestamp:    time.Now().Unix(),
				PDULength:    3,
				TrapIdentity: &TrapIdentityDTO{
					EnterpriseID: ".1.3.6.1.4.1.666.1",
					Generic:      6,
					Specific:     1,
				},
				Results: &SNMPResults{
					Results: []SNMPResultDTO{
						{
							Base: ".1.3.6.1.4.1.666.2.1.1",
							Value: SNMPValueDTO{
								Type:  4,
								Value: base64.StdEncoding.EncodeToString([]byte("This is a test")),
							},
						},
					},
				},
			},
		},
	}
	data, err := xml.Marshal(dto)
	assert.NilError(t, err)

	mock.Send(buildMessage("001", 0, 1, data))
	time.Sleep(1 * time.Second)
	assert.Assert(t, strings.Contains(message, "This is a test"))
	cli.Stop()
}

func TestNetflowParser(t *testing.T) {
	mock := &mockConsumer{
		msgChannel: make(chan *kafka.Message, 1),
	}
	cli := createKafkaClient(mock)
	cli.Parser = "netflow"
	var message string
	go func() {
		cli.Start(func(key, msg []byte) {
			fmt.Printf("Key %s, Value: %s\n", string(key), string(msg))
			message = string(msg)
		})
	}()

	ts := uint64(time.Now().Unix())
	netflow := &netflow.FlowMessage{
		NetflowVersion: netflow.NetflowVersion_V9,
		Direction:      netflow.Direction_EGRESS,
		Timestamp:      ts,
		DeltaSwitched:  &wrappers.UInt64Value{Value: ts},
		FirstSwitched:  &wrappers.UInt64Value{Value: ts},
		LastSwitched:   &wrappers.UInt64Value{Value: ts + 1},
		SrcAddress:     "11.0.0.1",
		DstAddress:     "12.0.0.2",
		NumBytes:       &wrappers.UInt64Value{Value: 1000},
	}
	netflowBytes, err := proto.Marshal(netflow)
	assert.NilError(t, err)

	location := "Test"
	systemID := "001"
	source := "10.0.0.1"
	port := uint32(8877)
	telemetryMsg := &telemetry.TelemetryMessageLog{
		Location:      &location,
		SystemId:      &systemID,
		SourceAddress: &source,
		SourcePort:    &port,
		Message: []*telemetry.TelemetryMessage{
			{
				Timestamp: &ts,
				Bytes:     netflowBytes,
			},
		},
	}
	data, err := proto.Marshal(telemetryMsg)
	assert.NilError(t, err)

	mock.Send(buildMessage("001", 0, 1, data))
	time.Sleep(1 * time.Second)
	assert.Assert(t, strings.Contains(message, "12.0.0.2"))
	cli.Stop()
}

func runProcessMessageTest(t *testing.T, wg *sync.WaitGroup, cli *KafkaClient, id string) {
	var data []byte
	data = cli.processMessage(buildMessage(id, 0, 3, []byte("ABC")))
	assert.Assert(t, data == nil)
	data = cli.processMessage(buildMessage(id, 1, 3, []byte("DEF")))
	assert.Assert(t, data == nil)
	data = cli.processMessage(buildMessage(id, 1, 3, []byte("DEF"))) // Make sure duplicates are handled properly
	assert.Assert(t, data == nil)
	data = cli.processMessage(buildMessage(id, 2, 3, []byte("GHI")))
	assert.Assert(t, data != nil)
	assert.Equal(t, "ABCDEFGHI", string(data))
	wg.Done()
}

func buildMessage(id string, chunk, total int32, data []byte) *kafka.Message {
	topic := "Test"
	sinkMsg := &sink.SinkMessage{
		MessageId:          id,
		CurrentChunkNumber: chunk,
		TotalChunks:        total,
		Content:            data,
	}
	bytes, _ := proto.Marshal(sinkMsg)
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Key:            []byte(id),
		Value:          bytes,
	}
	return kafkaMsg
}

func createKafkaClient(mock *mockConsumer) *KafkaClient {
	cli := &KafkaClient{
		Bootstrap: "127.0.0.1:9092",
		Topic:     "Test",
		GroupID:   "Test",
		IPC:       "sink",
		consumer:  mock,
	}
	cli.createVariables()
	cli.chunkProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mock_processed_chunk_total",
	})
	cli.msgProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mock_processed_messages_total",
	})
	return cli
}
