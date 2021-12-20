all: test build

jenkins: install_dep_tool install_go_linter test build

install_dep_tool:
	go get github.com/tools/godep

install_go_linter:
	go get -u -v golang.org/x/lint/golint

initialize_deps:
	go get -d -v ./...
	go get -d -v github.com/stretchr/testify/assert
	go get -d -v github.com/onsi/ginkgo
	go get -d -v github.com/onsi/gomega
	go get -v golang.org/x/lint/golint

update_deps:
	go install -v ./...
	go install -v github.com/stretchr/testify/assert
	go install -v github.com/onsi/ginkgo
	go install -v github.com/onsi/gomega
	go install -v golang.org/x/lint/golint

test:
	go vet ./...
	go test -v ./...

build:
	go build ./...


.PHONY: all install_dep_tool install_go_linter initialize_deps update_deps test jenkins
