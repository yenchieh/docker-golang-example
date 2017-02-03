#
# docker build --rm -t cacoo-site .
FROM golang:1.7
MAINTAINER Yen-Chieh Chen <yenchiehchen0830@gmail.com>

RUN mkdir -p /go/src/show-dev
WORKDIR /go/src/show-dev
ADD . /go/src/show-dev
RUN go build ./main.go
ENTRYPOINT "/go/src/show-dev/main"
EXPOSE 20001