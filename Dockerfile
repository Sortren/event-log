FROM golang:1.18

WORKDIR /event-log

COPY . .

RUN go build -o bin/server src/main.go

