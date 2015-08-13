all: deps build test

deps:
	go get -t -v ./...

test:
	cd clock; go test .; cd ..
	cd safemap; go test .; cd ..
	cd circuitbreaker; go test .; cd ..
	cd timetoken; go test .; cd ..
	cd retrier; go test .; cd ..

build:
	go vet
	go build ./...



.PHONY: deps test
