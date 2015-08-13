all: deps build test

deps:
	go get -t -v ./...

test:
	go test -v ./...

build:
	go vet
	go build ./...



.PHONY: deps test
