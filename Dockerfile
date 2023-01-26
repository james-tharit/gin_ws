FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .
ADD .air.toml .

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

ENTRYPOINT air -c .air.toml