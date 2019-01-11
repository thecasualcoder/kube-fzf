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

all: setup build

ensure-out-dir:
	mkdir -p out

modules: ## add missing and remove unused modules
	go mod tidy

compile: ensure-out-dir ## compiles kube-tmuxp for this platform
	$(GOBIN) build -ldflags "-X main.version=${BUILD}" -o  ./out/findpod ./findpod.go

compile-linux: ensure-out-dir ## compiles kube-tmuxp for linux
	GOOS=linux GOARCH=amd64 $(GOBIN) build -ldflags "-X main.version=${BUILD}" -o ./out/findpod ./findpod.go

fmt: ## format go code
	$(GOBIN) fmt $(SRC_PACKAGES)

vet: ## examine go code for suspicious constructs
	$(GOBIN) vet $(SRC_PACKAGES)

setup: ## setup environment
ifeq ($(RICHGO),)
	$(GOBIN) get -u github.com/kyoh86/richgo
endif

build: fmt vet compile ## build the application

