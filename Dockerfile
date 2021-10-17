# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /keyValueApp

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o github.com/cemezgn/keyValueApp

EXPOSE 8090

CMD [ "/keyValueApp" ]