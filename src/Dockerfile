FROM golang:1.17-alpine

WORKDIR /src
COPY . /src

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=auto

RUN go mod download

EXPOSE 8080
