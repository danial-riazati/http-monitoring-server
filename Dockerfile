FROM golang:1.18-buster

MAINTAINER danial riazati <dnr1380@gmail.com>

RUN mkdir /app

ADD . /app

WORKDIR /app
RUN go get -d -v
RUN go build -v



CMD ["./http-monitoring-server"]