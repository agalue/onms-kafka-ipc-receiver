FROM golang:alpine AS builder
RUN mkdir /app && \
    echo "@edgecommunity http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    apk update && \
    apk add --no-cache alpine-sdk git librdkafka-dev@edgecommunity
ADD ./ /app/
WORKDIR /app
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -a -o sink-receiver

FROM alpine
RUN echo "@edgecommunity http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    apk update && \
    apk add --no-cache bash librdkafka@edgecommunity && \
    rm -rf /var/cache/apk/* && \
    addgroup -S onms && adduser -S -G onms onms
COPY --from=builder /app/sink-receiver /sink-receiver
COPY ./docker-entrypoint.sh /
USER onms
LABEL maintainer="Alejandro Galue <agalue@opennms.org>" \
      name="OpenNMS Sink API Receiver"
ENTRYPOINT [ "/docker-entrypoint.sh" ]
