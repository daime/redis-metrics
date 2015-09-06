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
	go test -v ./...

clean:
	go clean

docker:
	docker-compose rm -f
	docker-compose build
	docker-compose up
