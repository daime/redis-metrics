all: install

run: install
	redis-metrics

install: dependencies
	go install github.com/daime/redis-metrics

build: dependencies
	go build ./...

dependencies:
	go get ./...

test:
	go test ./...

clean:
	go clean
