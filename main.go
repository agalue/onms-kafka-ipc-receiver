// A sample kafka consumer that works with single or multi-part messages
// There are multiple ways to implement this, so use this as a reference only.
//
// @author Alejandro Galue <agalue@opennms.org>

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/agalue/sink-receiver/client"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// The main function
func main() {
	cli := client.KafkaClient{}
	flag.StringVar(&cli.Bootstrap, "bootstrap", "localhost:9092", "kafka bootstrap server")
	flag.StringVar(&cli.Topic, "topic", "OpenNMS.Sink.Trap", "kafka topic that will receive the messages")
	flag.StringVar(&cli.GroupID, "group-id", "sink-go-client", "the consumer group ID")
	flag.Var(&cli.Parameters, "parameter", "Kafka consumer configuration attribute (can be used multiple times)\nfor instance: acks=1")
	flag.BoolVar(&cli.IsFlow, "is-flow", false, "Set to true if the payload is a Flow message")
	flag.Parse()

	log.Println("starting consumer")
	if err := cli.Initialize(); err != nil {
		panic(err)
	}
	log.Println("consumer started")

	go cli.Start(func(msg []byte) {
		/////////////////////////////////////////////
		// TODO Implement your custom actions here //
		/////////////////////////////////////////////

		log.Printf("message received: %s", string(msg))
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
