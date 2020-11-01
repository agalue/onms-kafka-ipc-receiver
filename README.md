OpenNMS Kafka IPC Receiver [![Go Report Card](https://goreportcard.com/badge/github.com/agalue/onms-kafka-ipc-receiver)](https://goreportcard.com/report/github.com/agalue/onms-kafka-ipc-receiver)
====

A sample Kafka Consumer application to display IPC messages into the standard output for troubleshooting purposes. It supports reconstructing split messages when the payload exceeds the limit.

It exposes Prometheus compatible metrics through port 8181, using the `/metrics` endpoint.

This repository also contains a Dockerfile to compile and build an image with the tool, which can be fully customized through environment variables.

Inside the `protobuf` directory, the `.proto` files extracted from OpenNMS source code contain the Protobuf definitions. If those files change in OpenNMS, make sure to re-generate the protobuf code by using the [build.sh](protobuf/build.sh) command, which expects to have `protoc` installed on your system.

## Requirements

When using Docker:

* `BOOTSTRAP_SERVER` environment variable with Kafka Bootstrap Server (i.e. `kafka01:9092`)
* `IPC` the IPC message kind to process. Either `rpc` or `sink` is allowed (defaults to `sink`).
* `TOPIC` environment variable with the source Sink API Kafka Topic with GPB Payload.
* `PARSER` the parser to use when processing Sink Messages. Valid values are: `syslog`, `netflow`, `flowdocument` (used only when `IPC=sink`).
* `GROUP_ID` environment variable with the Consumer Group ID (defaults to `opennms`)
* To pass consumer settings, add an environment variable with the prefix `KAFKA_`, for example: `KAFKA_AUTO_OFFSET_RESET`.

For consumer settings, the character underscore will be replaced with a dot and converted to lowercase. For example, `KAFKA_AUTO_OFFSET_RESET` will be configured as `auto.offset.reset`.

When using CLI:

Use `--help` for more details.

## Build

In order to build the application:

```bash
docker build -t agalue/onms-kafka-ipc-receiver:latest .
docker push agalue/onms-kafka-ipc-receiver:latest
```

> *NOTE*: Please use your own Docker Hub account or use the image provided on my account.

To build the controller locally for testing:

```bash
export GO111MODULE="on"

go build
./onms-kafka-ipc-receiver
```
