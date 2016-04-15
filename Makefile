all: deps build test

deps:
	go get -t ./...

test:
	go test ./...

build:
	go vet ./...
	go build ./...



.PHONY: deps test
