OpenNMS Kafka IPC Receiver [![Go Report Card](https://goreportcard.com/badge/github.com/agalue/onms-kafka-ipc-receiver)](https://goreportcard.com/report/github.com/agalue/onms-kafka-ipc-receiver)
====

A sample Kafka Consumer application to display IPC messages into the standard output for troubleshooting purposes. It supports reconstructing split messages when the payload exceeds the limit.

It exposes Prometheus compatible metrics through port 8181, using the `/metrics` endpoint.

This repository also contains a `Dockerfile` to compile and build a Docker Image with the tool, which can be fully customized through environment variables.

Inside the `protobuf` directory, the `.proto` files extracted from OpenNMS source code contain the Protobuf definitions. If those files change in OpenNMS, make sure to re-generate the protobuf code by using the [build.sh](protobuf/build.sh) command, which expects to have `protoc` installed on your system.

> This has been only tested against Horizon 27 and Meridian 2020.

## Usage

When using Docker:

* `BOOTSTRAP_SERVER` environment variable with Kafka Bootstrap Server (i.e. `kafka01:9092`)
* `IPC` the IPC message kind to process. Either `rpc` or `sink` is allowed (defaults to `sink`).
* `TOPIC` environment variable with the source Sink API Kafka Topic with GPB Payload.
* `PARSER` the parser to use when processing Sink Messages. Valid values are: `syslog`, `snmp`, `netflow`.
* `GROUP_ID` environment variable with the Consumer Group ID (defaults to `opennms`)
* To pass consumer settings, add an environment variable with the prefix `KAFKA_`, for example: `KAFKA_AUTO_OFFSET_RESET`.

For consumer settings, the character underscore will be replaced with a dot and converted to lowercase. For example, `KAFKA_AUTO_OFFSET_RESET` will be configured as `auto.offset.reset`.

When using CLI:

Use `--help` for more details.

## Build

To build the application using Docker:

```bash
docker build -t agalue/onms-kafka-ipc-receiver:latest .
docker push agalue/onms-kafka-ipc-receiver:latest
```

> *NOTE*: Please use your own Docker Hub account or use the image provided on my account.

To build the applicatoin locally, make sure you have [Go](https://golang.org/) 1.16 installed on your machine, as well as [librdkafka](https://github.com/edenhill/librdkafka);then:

```bash
go build
```

## Sample Output

### Syslog (Sink API)

To run the parser:

```bash
onms-kafka-ipc-receiver -bootstrap kafka:9092 -ipc sink -parser syslog -topic OpenNMS.Sink.Syslog
```

The Protobuf payload is parsed and the tool prints a human-readable representation of it in JSON:

```json
```

### SNMP Traps (Sink API)

To run the parser:

```bash
onms-kafka-ipc-receiver -bootstrap kafka:9092 -ipc sink -parser snmp -topic OpenNMS.Sink.Trap
```

The Protobuf payload is parsed and the tool prints a human-readable representation of it in JSON:

```json
{
  "location": "Apex",
  "systemId": "minion01",
  "trapAddress": "172.18.0.1",
  "messages": [
    {
      "agentAddress": "172.18.0.1",
      "community": "public",
      "version": "v2",
      "timestamp": 22600476,
      "creationTime": 1615473381643,
      "pduLength": 6,
      "trapIdentity": {
        "enterpriseID": ".1.3.6.1.4.1.9.9.171.2",
        "generic": 6,
        "specific": 2
      },
      "results": {
        "varbinds": [
          {
            "base": ".1.3.6.1.6.3.18.1.3.0",
            "value": {
              "type": 64,
              "value": "11.0.0.4"
            }
          },
          {
            "base": ".1.3.6.1.4.1.9.9.171.1.2.2.1.6",
            "value": {
              "type": 4,
              "value": "0xd047b002"
            }
          },
          {
            "base": ".1.3.6.1.4.1.9.9.171.1.2.2.1.7",
            "value": {
              "type": 4,
              "value": "0xd047b002"
            }
          },
          {
            "base": ".1.3.6.1.4.1.9.9.171.1.2.3.1.16",
            "value": {
              "type": 2,
              "value": "10"
            }
          }
        ]
      }
    }
  ]
}
```

### Flows (Sink API)

To run the parser:

```bash
onms-kafka-ipc-receiver -bootstrap kafka:9092 -ipc sink -parser netflow -topic OpenNMS.Sink.Telemetry-Netflow-9
```

The Protobuf payload is parsed and the tool prints a human-readable representation of it in JSON:

```json

```

### RPC

To run the parser for requests (assuming `single-topic` is enabled in OpenNMS and Minion):

```bash
onms-kafka-ipc-receiver -bootstrap kafka:9092 -ipc rpc -topic OpenNMS.Apex.rpc-request
```

The above assumes the location of where the Minions live is called `Apex`.

To run the parser for responses (assuming `single-topic` is enabled in OpenNMS and Minion):

```bash
onms-kafka-ipc-receiver -bootstrap kafka:9092 -ipc rpc -topic OpenNMS.rpc-response
```

The Protobuf payload is parsed and the tool prints a human-readable representation of it in XML:

```
```
