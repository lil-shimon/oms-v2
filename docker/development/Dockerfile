FROM golang:1.15-alpine

WORKDIR /go/src/server
COPY ./server .

RUN apk update && \
  apk add --no-cache git

EXPOSE 1323
