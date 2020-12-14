FROM golang:alpine3.12

RUN export GOPATH=/var/golang       \
    && export GOROOT=/go            \
    && apk update                   \
    && apk --no-cache add           \
    gcc                             \
    git                             \
    musl-dev                        \
    npm                             \
    postgresql                      \
    vim                             \
    && npm install -g               \
    typescript
# gcc: go test
# musl-dev: standard library for go test

ENV GOROOT /usr/local/go
ENV GOPATH /var/golang

WORKDIR /var/golang/src/