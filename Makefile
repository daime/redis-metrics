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

coverage-ci:
	mkdir -p _test
	go get -u github.com/pierrre/gotestcover
	gotestcover -coverprofile=_test/cover.out ./...
	go tool cover -html=_test/cover.out -o=_test/cover.html

coverage: coverage-ci
	open _test/cover.html

clean:
	go clean

docker-clean:
	docker-compose rm -f

docker: build docker-clean
	docker-compose build
	docker-compose up
