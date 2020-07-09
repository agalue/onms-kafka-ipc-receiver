OpenNMS Sink API Receiver
====

A sample Kafka Consumer application to display sink into the standard output for troubleshooting purposes. It supports reconstructing split messages when the payload exceeds the limit.

This repository also contains a Dockerfile to compile and build an image with the tool, which can be fully customized through environment variables, so the solution can be used with Kubernetes (the sample YAML file is also available).

Inside the `protobuf` directory, the [sink.proto](protobuf/sink.proto) file extracted from OpenNMS source code contains the Protobuf definitions. If that file changes, make sure to re-generate the protobuf code by using the [build.sh](protobuf/build.sh) command, which expects to have `protoc` installed on your system.

## Requirements

* `BOOTSTRAP_SERVERS` environment variable with Kafka Bootstrap Server (i.e. `kafka01:9092`)
* `TOPIC` environment variable with the source Sink API Kafka Topic with GPB Payload (defaults to `OpenNMS.Sink.Trap`)
* `GROUP_ID` environment variable with the Consumer Group ID (defaults to `opennms`)
* To pass consumer settings, add an environment variable with the prefix `KAFKA_`, for example: `KAFKA_AUTO_OFFSET_RESET`.

For consumer settings, the character underscore will be replaced with a dot and converted to lowercase. For example, `KAFKA_AUTO_OFFSET_RESET` will be configured as `auto.offset.reset`.

## Build

In order to build the application:

```bash
docker build -t agalue/sink-receiver-go:latest .
docker push agalue/sink-receiver-go:latest
```

> *NOTE*: Please use your own Docker Hub account or use the image provided on my account.

To build the controller locally for testing:

```bash
export GO111MODULE="on"

go build
./sink-receiver
```
