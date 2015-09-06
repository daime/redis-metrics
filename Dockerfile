FROM golang

ADD . /go/src/github.com/daime/redis-metrics

WORKDIR /go/src/github.com/daime/redis-metrics
ENTRYPOINT make run
