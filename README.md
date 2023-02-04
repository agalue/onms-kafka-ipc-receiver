OpenNMS Kafka IPC Receiver [![Go Report Card](https://goreportcard.com/badge/github.com/agalue/onms-kafka-ipc-receiver)](https://goreportcard.com/report/github.com/agalue/onms-kafka-ipc-receiver)
====

A sample Kafka Consumer application based on [Watermil](https://watermill.io/), powered by [Sarama](https://github.com/Shopify/sarama), written in [Go](https://golang.org/) to display IPC messages into the standard output for troubleshooting purposes. It supports reconstructing split messages when the payload exceeds the limit.

It exposes Prometheus compatible metrics through port 8181, using the `/metrics` endpoint.

This repository also contains a `Dockerfile` to compile and build a Docker Image with the tool, which can be fully customized through environment variables.

Inside the `protobuf` directory, the `.proto` files extracted from OpenNMS source code contain the Protobuf definitions. If those files change in OpenNMS, make sure to re-generate the protobuf code by using the [build.sh](protobuf/build.sh) command, which expects to have `protoc` installed on your system.

> This has been only tested against Horizon 27 and Meridian 2020.

## Usage

When using CLI or Docker:

Use `--help` for more details.

## Build

To build the application using Docker:

```bash
docker build -t agalue/onms-kafka-ipc-receiver:latest .
docker push agalue/onms-kafka-ipc-receiver:latest
```

> *NOTE*: Please use your own Docker Hub account or use the image provided on my account.

To build the applicatoin locally, make sure you have Go 1.16 installed on your machine, then:

```bash
go build
```

## Sample Output

### Heartbeat (Sink API)

To run the parser:

```bash
onms-kafka-ipc-receiver -bootstrap kafka:9092 -ipc sink -parser heartbeat -topic OpenNMS.Sink.Heartbeat
```

The Protobuf payload is parsed and the tool prints a human-readable representation of it in XML:

```xml
<minion>
  <id>minion01</id>
  <location>Apex</location>
  <timestamp>2021-03-26T14:42:53.803-04:00</timestamp>
</minion>
```

### Syslog (Sink API)

To run the parser:

```bash
onms-kafka-ipc-receiver -bootstrap kafka:9092 -ipc sink -parser syslog -topic OpenNMS.Sink.Syslog
```

The Protobuf payload is parsed and the tool prints a human-readable representation of it in JSON:

```json
{
  "systemId": "minion01",
  "location": "Apex",
  "sourceAddress": "192.168.75.1",
  "sourcePort": 60209,
  "messages": [
    {
      "timestamp": "2021-03-26T14:49:27.734-04:00",
      "content": "\u003c190\u003e2021-03-26T14:49:27-04:00 agalue-mbp.local udpgen[9601]: %%SEC-6-IPACCESSLOGP: list in110 denied tcp 10.99.99.1(63923) -\u003e 10.98.98.1(1521), 1 packet"
    }
  ]
}
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
onms-kafka-ipc-receiver -bootstrap kafka:9092 -ipc sink -parser netflow -topic OpenNMS.Sink.Telemetry-Netflow-5
```

The Protobuf payload is parsed and the tool prints a human-readable representation of it in JSON:

```json
{
  "timestamp": 1616785647091,
  "num_bytes": {
    "value": 295
  },
  "dst_address": "47.239.92.73",
  "dst_as": {},
  "dst_mask_len": {
    "value": 29
  },
  "dst_port": {
    "value": 32988
  },
  "engine_id": {},
  "engine_type": {
    "value": 1
  },
  "first_switched": {
    "value": 1621080613225
  },
  "last_switched": {
    "value": 1621080613311
  },
  "num_flow_records": {
    "value": 8
  },
  "num_packets": {
    "value": 577
  },
  "flow_seq_num": {
    "value": 1
  },
  "input_snmp_ifindex": {
    "value": 50575
  },
  "output_snmp_ifindex": {
    "value": 43523
  },
  "next_hop_address": "61.36.170.15",
  "protocol": {
    "value": 6
  },
  "sampling_interval": {},
  "src_address": "137.8.59.230",
  "src_as": {},
  "src_mask_len": {
    "value": 27
  },
  "src_port": {
    "value": 54717
  },
  "tcp_flags": {},
  "tos": {}
}
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

In both cases, the Protobuf payload is parsed and the tool prints a human-readable representation of it in XML.