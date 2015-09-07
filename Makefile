all: install

run: install
	redis-metrics

install: dependencies
	go install ./...

build: dependencies
	go build ./...

dependencies:
	go get ./...

test:
	go test -v ./...

clean:
	go clean

docker-clean:
	docker-compose rm -f

docker: docker-clean
	docker-compose build
	docker-compose up
