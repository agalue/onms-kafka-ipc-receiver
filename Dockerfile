FROM golang:alpine AS builder
RUN mkdir /app && \
    apk update && \
    apk add --no-cache alpine-sdk git
ADD ./ /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags musl -a -o onms-kafka-ipc-receiver

FROM alpine
RUN apk update && \
    apk add --no-cache bash && \
    rm -rf /var/cache/apk/* && \
    addgroup -S onms && adduser -S -G onms onms
COPY --from=builder /app/onms-kafka-ipc-receiver /onms-kafka-ipc-receiver
COPY ./docker-entrypoint.sh /
USER onms
LABEL maintainer="Alejandro Galue <agalue@opennms.org>" \
      name="OpenNMS Kafka IPC API Receiver"
ENTRYPOINT [ "/docker-entrypoint.sh" ]
