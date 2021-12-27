all: clean build test

update_dep:
	go get $(DEP)
	go mod tidy

update_all_deps:
	go get -u
	go mod tidy

test:
	go vet ./...
	go clean -testcache
	go test -v ./...

build:
	go build -o examples/crontask examples/crontask/example.go
	go build -o examples/scheduledtask examples/scheduledtask/example.go

clean:
	rm -f examples/scheduledtask/example
	rm -f examples/crontask/example


.PHONY: all update_dep update_all_deps test build clean
