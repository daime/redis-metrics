all: build

run: build
	./bin/redis-metrics

build: dependencies
	go build .

dependencies:
	go get -u -d ./...

test:
	docker-compose up
