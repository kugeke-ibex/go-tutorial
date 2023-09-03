FROM golang:1.21-rc-alpine

WORKDIR /go/src/app/
COPY ./goapp .

RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev
