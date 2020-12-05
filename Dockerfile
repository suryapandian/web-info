FROM golang:1.15-alpine AS builder

ADD . /go/src/web-info

WORKDIR /go/src/web-info

RUN go build -mod=vendor -o web-info .

EXPOSE 3001

ENTRYPOINT [ "./web-info"]