// A sample kafka consumer that works with single or multi-part messages
// There are multiple ways to implement this, so use this as a reference only.
//
// @author Alejandro Galue <agalue@opennms.org>

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/agalue/onms-kafka-ipc-receiver/client"
	"github.com/agalue/onms-kafka-ipc-receiver/protobuf/netflow"
	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// The main function
func main() {
	var isNetflow bool
	cli := client.KafkaClient{}
	flag.StringVar(&cli.Bootstrap, "bootstrap", "localhost:9092", "kafka bootstrap server")
	flag.StringVar(&cli.Topic, "topic", "OpenNMS.Sink.Trap", "kafka topic that will receive the messages")
	flag.StringVar(&cli.GroupID, "group-id", "sink-go-client", "the consumer group ID")
	flag.Var(&cli.Parameters, "parameter", "Kafka consumer configuration attribute (can be used multiple times)\nfor instance: acks=1")
	flag.StringVar(&cli.IPC, "ipc", "sink", "IPC API, either 'sink' or 'rpc'")
	flag.BoolVar(&cli.IsTelemetry, "is-telemetry", false, "Set to true if the payload is a telemetry message")
	flag.BoolVar(&isNetflow, "is-netflow", false, "Set to true if the telemetry payload is a netflow packet")
	flag.Parse()

	log.Println("starting consumer")
	if err := cli.Initialize(); err != nil {
		log.Fatalf("Cannot initialize consumer: %v", err)
	}
	log.Println("consumer started")

	go cli.Start(func(msg []byte) {
		/////////////////////////////////////////////
		// TODO Implement your custom actions here //
		/////////////////////////////////////////////

		log.Println("message received")
		if cli.IsTelemetry && isNetflow {
			flow := &netflow.FlowMessage{}
			err := proto.Unmarshal(msg, flow)
			if err != nil {
				log.Printf("warning: invalid netflow message received: %v", err)
				return
			}
			bytes, _ := json.MarshalIndent(flow, "", "  ")
			fmt.Println(string(bytes))
		} else {
			fmt.Println(string(msg))
		}
	})

	go func() {
		port := 8181
		log.Printf("Starting Prometheus Metrics Server on port %d", port)
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	cli.Stop()
}
