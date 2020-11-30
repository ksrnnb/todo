FROM golang:alpine3.12

RUN export GOPATH=/var/golang       \
    && export GOROOT=/go            \
    && apk update                   \
    && apk add                      \
    git                             \
    postgresql                      \
    vim

ENV GOROOT /usr/local/go
ENV GOPATH /var/golang

WORKDIR /var/golang/src/