EXECUTABLE = truffls
TOOLS_MOD_DIR = ./tools
TOOLS_DIR = $(abspath ./.tools)

GO_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

VERSION ?= $(shell git describe --tags 2>/dev/null || git rev-parse --short HEAD)
BUILD_DATE ?= $(shell date "+%Y-%m-%d")

GO_LDFLAGS := -X github.com/foodarchive/$(EXECUTABLE)/internal/config.BuildDate=$(BUILD_DATE) \
	-X github.com/foodarchive/$(EXECUTABLE)/internal/config.Version=$(VERSION)

GO_TEST_MIN = go test -v -timeout 30s
GO_TEST = $(GO_TEST_MIN) -race
GO_TEST_WITH_COVERAGE = $(GO_TEST) -coverprofile=coverage.txt -covermode=atomic

.DEFAULT_GOAL = help

$(TOOLS_DIR)/golangci-lint: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: build
build: $(GO_FILES) ## build the application
	@go build -trimpath -ldflags "$(GO_LDFLAGS)" -o ./bin/$(EXECUTABLE) ./cmd/$(EXECUTABLE)

.PHONY: test
test: ## run the test cases
	$(GO_TEST) ./...

.PHONY: test-with-coverage
test-with-coverage: ## run testing with code coverage and output html file
	$(GO_TEST_WITH_COVERAGE) ./... && go tool cover -html=coverage.txt -o coverage.html

.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint ## lint the source code using golangci
	$(TOOLS_DIR)/golangci-lint run --fix && $(TOOLS_DIR)/golangci-lint run

.PHONY: check-clean-work-tree
check-clean-work-tree: ## check whether the working branch is clean or not
	@if ! git diff --quiet; then \
	  echo; \
	  echo 'Working tree is not clean, did you forget to run "make precommit"?'; \
	  echo; \
	  git status; \
	  exit 1; \
	fi

.PHONY: check-license
check-license: ## make sure all go files have license header
	@licRes=$$(for f in $(GO_FILES) ; do \
	           awk '/Copyright The Truffls Contributors.|generated|GENERATED/ && NR<=3 { found=1; next } END { if (!found) print FILENAME }' $$f; \
	   done); \
	   if [ -n "$${licRes}" ]; then \
	           echo "license header checking failed:"; echo "$${licRes}"; \
	           exit 1; \
	   fi

.PHONY: ci
ci: check-clean-work-tree check-license lint test-with-coverage ## list of task that going to be executed on ci

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## display the version of the app
	@echo $(VERSION)

