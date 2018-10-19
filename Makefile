all: test build

jenkins: install_dep_tool install_go_linter production_restore_deps test build

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
	godep save ./...

update_deps:
	godep get -d -v ./...
	godep get -d -v github.com/stretchr/testify/assert
	godep get -d -v github.com/onsi/ginkgo
	godep get -d -v github.com/onsi/gomega
	godep get -v golang.org/x/lint/golint
	godep update ./...

test:
	golint ./...
	godep go vet ./...
	godep go test -v ./...

build:
	godep go build ./...

production_restore_deps:
	godep restore

.PHONY: all install_dep_tool install_go_linter initialize_deps update_deps test jenkins
