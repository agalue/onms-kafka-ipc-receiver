// A sample kafka consumer that works with single or multi-part messages
// There are multiple ways to implement this, so use this as a reference only.
//
// @author Alejandro Galue <agalue@opennms.org>

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/agalue/onms-kafka-ipc-receiver/client"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetOutput(os.Stdout)
	promPort := 8181

	cli := client.KafkaClient{}
	flag.StringVar(&cli.Bootstrap, "bootstrap", "localhost:9092", "kafka bootstrap server")
	flag.StringVar(&cli.Topic, "topic", "OpenNMS.Sink.Trap", "kafka topic that will receive the messages")
	flag.StringVar(&cli.GroupID, "group-id", "sink-go-client", "the consumer group ID")
	flag.StringVar(&cli.IPC, "ipc", "sink", "IPC API: sink, rpc")
	flag.StringVar(&cli.Parser, "parser", "snmp", "Sink API Parser: "+client.AvailableParsers.EnumAsString())
	flag.IntVar(&promPort, "prometheus-port", promPort, "Port to export Prometheus metrics")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
	}()
	go func() {
		select {
		case <-signalChan:
			cancel()
		case <-ctx.Done():
		}
	}()

	if err := cli.Initialize(ctx); err != nil {
		log.Fatalf("cannot initialize consumer: %v", err)
	}

	go func() {
		log.Printf("starting Prometheus Metrics Server on port %d", promPort)
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(fmt.Sprintf(":%d", promPort), nil)
	}()

	log.Println("starting consumer")
	cli.Start(func(msg []byte) {
		log.Printf("received %s:%s message: %s", cli.IPC, cli.Parser, string(msg))
	})
}
