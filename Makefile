all: deps build test

deps:
	go get -t ./...
	go get -u github.com/golang/lint/golint

test:
	go test ./...

build:
	golint ./...
	go vet ./...
	go build ./...



.PHONY: deps test
