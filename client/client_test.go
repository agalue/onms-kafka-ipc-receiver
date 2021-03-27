// @author Alejandro Galue <agalue@opennms.org>

package client

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/netflow"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/sink"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/telemetry"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/prometheus/client_golang/prometheus"
	"gotest.tools/assert"
)

func TestProcessSingleMessage(t *testing.T) {
	cli, _, cancel := createKafkaClient()
	data := cli.processMessage(buildMessage("0001", 0, 1, []byte("ABC")))
	assert.Equal(t, "ABC", string(data))
	cancel()
}

func TestProcessMultipleMessages(t *testing.T) {
	cli, _, cancel := createKafkaClient()
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go runProcessMessageTest(t, wg, cli, fmt.Sprintf("ID%d", i))
	}
	wg.Wait()
	cancel()
}

func TestSyslogParser(t *testing.T) {
	cli, sub, cancel := createKafkaClient()

	cli.Parser = "syslog"
	var message string
	go func() {
		cli.Start(func(key string, msg []byte) {
			fmt.Printf("Key %s, Value: %s\n", key, string(msg))
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

	sub.Publish("Test", buildMessage("001", 0, 1, data))
	time.Sleep(1 * time.Second)
	assert.Assert(t, strings.Contains(message, "This is a test"))
	cancel()
}

func TestSnmpParser(t *testing.T) {
	cli, sub, cancel := createKafkaClient()
	cli.Parser = "snmp"
	var message string
	go func() {
		cli.Start(func(key string, msg []byte) {
			fmt.Printf("Key %s, Value: %s\n", key, string(msg))
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

	sub.Publish("Test", buildMessage("001", 0, 1, data))
	time.Sleep(1 * time.Second)
	assert.Assert(t, strings.Contains(message, "This is a test"))
	cancel()
}

func TestNetflowParser(t *testing.T) {
	cli, sub, cancel := createKafkaClient()
	cli.Parser = "netflow"
	var message string
	go func() {
		cli.Start(func(key string, msg []byte) {
			fmt.Printf("Key %s, Value: %s\n", key, string(msg))
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

	sub.Publish("Test", buildMessage("001", 0, 1, data))
	time.Sleep(1 * time.Second)
	assert.Assert(t, strings.Contains(message, "12.0.0.2"))
	cancel()
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

func buildMessage(id string, chunk, total int32, data []byte) *message.Message {
	sinkMsg := &sink.SinkMessage{
		MessageId:          id,
		CurrentChunkNumber: chunk,
		TotalChunks:        total,
		Content:            data,
	}
	bytes, _ := proto.Marshal(sinkMsg)
	return &message.Message{
		UUID:    id,
		Payload: bytes,
	}
}

func createKafkaClient() (*KafkaClient, *gochannel.GoChannel, context.CancelFunc) {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	ctx, cancel := context.WithCancel(context.Background())
	msgChannel, _ := pubSub.Subscribe(ctx, "Test")

	cli := &KafkaClient{
		Bootstrap:  "127.0.0.1:9092",
		Topic:      "Test",
		GroupID:    "Test",
		IPC:        "sink",
		msgChannel: msgChannel,
		subscriber: &kafka.Subscriber{},
	}
	cli.createVariables()
	cli.chunkProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mock_processed_chunk_total",
	})
	cli.msgProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mock_processed_messages_total",
	})
	return cli, pubSub, cancel
}
