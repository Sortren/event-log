FROM golang:1.18

WORKDIR /event-log

COPY . .

RUN go mod download && go mod tidy

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /event-log/src/cmd

RUN swag init --parseDependency --parseInternal

RUN go build -o /bin/server main.go
