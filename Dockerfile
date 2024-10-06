FROM golang:latest

WORKDIR /go/app

RUN apt-get update && apt install librdkafka-dev

RUN apk --no-cache update && \
    apk --no-cache add git gcc libc-dev

# Kafka Go client is based on the C library librdkafka
ENV CGO_ENABLED 1
ENV GOFLAGS -mod=vendor
ENV GOOS=linux
ENV GOARCH=amd64

RUN export GO111MODULE=on

CMD [ "tail", "-f", "/dev/null" ]