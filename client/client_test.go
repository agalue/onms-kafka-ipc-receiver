// @author Alejandro Galue <agalue@opennms.org>

package client

import (
	"fmt"
	"sync"
	"testing"

	"github.com/agalue/sink-receiver/protobuf/sink"
	"github.com/golang/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"gotest.tools/assert"
)

func TestCreateConfig(t *testing.T) {
	cli := &KafkaClient{
		Bootstrap:  "kafka1:9092",
		GroupID:    "GoTest",
		Parameters: "acks=1,fetch.min.bytes=2048,ssl.truststore.location",
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
	cli := &KafkaClient{}
	cli.createVariables()
	data := cli.processMessage(buildMessage("0001", 0, 1, []byte("ABC")))
	assert.Equal(t, "ABC", string(data))
}

func TestProcessMultipleMessages(t *testing.T) {
	cli := &KafkaClient{}
	cli.createVariables()
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go runProcessMessageTest(t, wg, cli, fmt.Sprintf("ID%d", i))
	}
	wg.Wait()
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
