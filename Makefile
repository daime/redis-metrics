all: install

run: install
	redis-metrics

build: dependencies
	go install github.com/daime/redis-metrics

dependencies:
	go get github.com/garyburd/redigo/redis

test:
	go test

clean:
	go clean

env:
	export GOPATH=${GOPATH}:${PWD}
