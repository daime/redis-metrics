WATCH_INTERVAL=5
TARGET=dist/

all: install

install: dependencies
	go install ./...

build: dependencies
	go build ./...

dist: dependencies
	mkdir -p ${TARGET}
	GOARCH=amd64 GOOS=linux go build -o ${TARGET}redis-metrics-linux-amd64
	GOARCH=386 GOOS=linux go build -o ${TARGET}redis-metrics-linux-386
	GOARCH=amd64 GOOS=darwin go build -o ${TARGET}redis-metrics-darwin-amd64

release: dist
	./ci/create-release

dependencies:
	go get ./...

test:
	go test ./...
	go test -race ./...

watch:
	watch -n ${WATCH_INTERVAL} make clean test

coverage-ci:
	mkdir -p _test
	go get -u github.com/pierrre/gotestcover
	gotestcover -coverprofile=_test/cover.out ./...
	go tool cover -html=_test/cover.out -o=_test/cover.html

coverage: coverage-ci
	open _test/cover.html

clean:
	rm -rf ${TARGET}
	go clean

docker-clean:
	docker-compose rm -f

docker: build docker-clean
	docker-compose build
	docker-compose up
