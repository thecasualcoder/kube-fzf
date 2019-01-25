.PHONY: help
help: ## prints help (only for tasks with comment)
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

SRC_PACKAGES=$(shell go list ./...)
BUILD?=$(shell git describe --tags --always --dirty)
RICHGO=$(shell command -v richgo 2> /dev/null)

ifeq ($(RICHGO),)
	GOBIN=go
else
	GOBIN=richgo
endif

all: setup fmt vet tests example

ensure-out-dir:
	mkdir -p out

modules: ## add missing and remove unused modules
	go mod tidy -v

fmt: ## format go code
	$(GOBIN) fmt $(SRC_PACKAGES)

vet: ## examine go code for suspicious constructs
	$(GOBIN) vet $(SRC_PACKAGES)

setup: ## setup environment
ifeq ($(RICHGO),)
	$(GOBIN) get -u github.com/kyoh86/richgo
endif

mocks: ## generate mocks for testing
	./scripts/mocks

tests: ensure-out-dir # run all tests
	$(GOBIN) test $(SRC_PACKAGES) -p=1 -coverprofile ./out/coverage -v

tests-cover-html: ensure-out-dir ## run all tests and generates coverage report in html
	@echo "mode: count" > out/coverage-all.out
	$(foreach pkg, $(SRC_PACKAGES),\
	$(GOBIN) test -coverprofile=out/coverage.out -covermode=count $(pkg);\
	tail -n +2 out/coverage.out >> out/coverage-all.out;)
	$(GOBIN) tool cover -html=out/coverage-all.out -o out/coverage.html

example: ## run the example code
	$(GOBIN) run -v examples/main.go
