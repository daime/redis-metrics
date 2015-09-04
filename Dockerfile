FROM golang

ADD . /go/src/github.com/daime/redis-metrics

RUN go get github.com/garyburd/redigo/redis
RUN go install github.com/daime/redis-metrics

ENTRYPOINT /go/bin/redis-metrics
