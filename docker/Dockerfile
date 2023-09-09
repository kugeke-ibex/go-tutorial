FROM golang:1.21-rc-alpine

WORKDIR /go/src/Tutorial/

RUN apk upgrade --update && \
    apk --no-cache add bash git gcc musl-dev openssh && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest
