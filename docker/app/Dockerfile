FROM golang:1.21-rc-alpine

ARG username=devcontainer
ARG useruid=1000
ARG usergid=${useruid}

WORKDIR /go/src/Tutorial/

RUN \
    # apk
    apk upgrade --update && \
    apk --no-cache add \
    bash \
    git \
    gcc \
    musl-dev \
    openssh \
    zip \
    unzip \
    wget \
    sudo \
    && \
    
    # create user
    addgroup -g ${usergid} -S ${username} && \
    adduser -s /bin/bash -u ${useruid} -S ${username} -G ${username} && \
    echo ${username} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${username} && \
    chmod 0440 /etc/sudoers.d/${username}

ENV GOPATH=/home/${username}/go
ENV GO111MODULE on
USER ${username}
