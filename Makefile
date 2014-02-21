ERROR_COLOR=\033[31;01m
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

all: clean deps format build

build:
	go build -o client

clean:
	rm client

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

format:
	go fmt ./...

updatedeps:
	@echo "$(OK_COLOR)==> Updating all dependencies$(NO_COLOR)"
	@go get -d -v -u ./...
	@echo $(DEPS) | xargs -n1 go get -d -u

test: clean deps format
	@echo "$(OK_COLOR)==> Running tests...$(NO_COLOR)"
	@-go test -coverprofile=coverage.out github.com/geetarista/vindinium-starter-go/vindinium # -gocheck.v -gocheck.f ClientSuite -covermode=count
	@go tool cover -html=coverage.out -o coverage.html
	@if [ -x vindinium/vindinium.test ]; then rm vindinium/vindinium.test; fi

.PHONY: all clean deps format release test updatedeps
