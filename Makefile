EXECUTABLE = truffls
GO_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
BUILD_DATE ?= $(shell date "+%Y-%m-%d")

GO_LDFLAGS := -X github.com/foodarchive/$(EXECUTABLE)/internal/config.BuildDate=$(BUILD_DATE) \
	-X github.com/foodarchive/$(EXECUTABLE)/internal/config.Version=$(VERSION)

default: help

build: $(GO_FILES) ## build the application
	@go build -trimpath -ldflags "$(GO_LDFLAGS)" -o ./bin/$(EXECUTABLE) ./cmd/$(EXECUTABLE)

test: ## run the tests
	@go test ./...

help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

version: ## display the version of the app
	@echo $(VERSION)

