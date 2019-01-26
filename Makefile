.PHONY: help
help: ## prints help (only for tasks with comment)
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

SRC_PACKAGES=$(shell go list ./...)
BUILD?=$(shell git describe --tags --always --dirty)
DEP:=$(shell command -v dep 2> /dev/null)
RICHGO=$(shell command -v richgo 2> /dev/null)

ifeq ($(RICHGO),)
	GOBIN=go
else
	GOBIN=richgo
endif

all: setup build

ensure-out-dir:
	mkdir -p out

build-deps: ## install deps
	dep ensure -v

compile: ensure-out-dir ## compile for this platform
	$(GOBIN) build -ldflags "-X main.version=${BUILD}" -o  ./out/findpod ./cmd/findpod
	$(GOBIN) build -ldflags "-X main.version=${BUILD}" -o  ./out/describepod ./cmd/describepod

compile-linux: ensure-out-dir ## compile for linux
	GOOS=linux GOARCH=amd64 $(GOBIN) build -ldflags "-X main.version=${BUILD}" -o ./out/findpod ./cmd/findpod
	GOOS=linux GOARCH=amd64 $(GOBIN) build -ldflags "-X main.version=${BUILD}" -o ./out/describepod ./cmd/describepod

fmt: ## format go code
	$(GOBIN) fmt $(SRC_PACKAGES)

vet: ## examine go code for suspicious constructs
	$(GOBIN) vet $(SRC_PACKAGES)

setup: ## setup environment
ifeq ($(DEP),)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif

ifeq ($(RICHGO),)
	$(GOBIN) get -u github.com/kyoh86/richgo
endif

build: build-deps fmt vet compile ## build the application