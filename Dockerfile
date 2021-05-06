FROM golang:1.15-alpine3.12

RUN apk update && \
    apk upgrade && \
    apk add make git

RUN go get github.com/cespare/reflex
ENV CGO_ENABLED=0

WORKDIR /go/src/app

COPY . /go/src/app

CMD reflex  -s -r '\.go$$' go run app/http/main.go
