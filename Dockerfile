FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x
RUN apk update && apk add --no-cache alpine-sdk git
ADD ./ /app/
RUN go build -o onms-kafka-ipc-receiver .

FROM alpine
RUN apk update && \
    apk add --no-cache bash && \
    rm -rf /var/cache/apk/* && \
    addgroup -S onms && adduser -S -G onms onms
COPY --from=builder /app/onms-kafka-ipc-receiver /onms-kafka-ipc-receiver
USER onms
LABEL maintainer="Alejandro Galue <agalue@opennms.org>" \
      name="OpenNMS Kafka IPC API Receiver"
ENTRYPOINT [ "/onms-kafka-ipc-receiver" ]
