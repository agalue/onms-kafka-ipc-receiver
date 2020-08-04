// @author Alejandro Galue <agalue@opennms.org>

package client

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/agalue/sink-receiver/protobuf/sink"
	"github.com/golang/protobuf/proto"
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

func (mock *mockConsumer) CommitMessage(m *kafka.Message) ([]kafka.TopicPartition, error) {
	return []kafka.TopicPartition{}, nil
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

func TestClient(t *testing.T) {
	mock := &mockConsumer{
		msgChannel: make(chan *kafka.Message, 1),
	}
	cli := createKafkaClient(mock)
	var message string
	go func() {
		cli.Start(func(msg []byte) {
			fmt.Printf("%s\n", string(msg))
			message = string(msg)
		})
	}()
	mock.Send(buildMessage("001", 0, 1, []byte("This is a test")))
	time.Sleep(1 * time.Second)
	assert.Equal(t, "This is a test", message)
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
		MessageId:          &id,
		CurrentChunkNumber: &chunk,
		TotalChunks:        &total,
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
		consumer: mock,
	}
	cli.createVariables()
	cli.chunkProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "onms_sink_processed_messages_total",
	})
	cli.msgProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "onms_sink_processed_chunk_total",
	})
	return cli
}
