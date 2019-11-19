export GO111MODULE := on
TEST?=$$(go list ./...)
GOFMT_FILES?=$$(find . -name '*.go')

default: build

setup:
    curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.21.0
.PHONY: setup

test:
	@echo "==> Running tests..."
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4
.PHONY: test

lint:
	@echo "==> Linting all packages..."
	golangci-lint run ./... -E gofmt -E golint
.PHONY: lint

fmt:
	gofmt -s -w $(GOFMT_FILES)
.PHONY: fmt

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./package-name"; \
		exit 1; \
	fi
	@echo "==> Compiling test binary..."
	go test -c $(TEST) $(TESTARGS)
.PHONY: test-compile

# Build targets
clean:
	@echo "==> Cleaning build artifacts..."
	go clean ./...
.PHONY: clean

gitsha := $(shell git log -n1 --pretty='%h')
version=$(shell git describe --exact-match --tags "$(gitsha)" 2>/dev/null)
ifeq ($(version),)
	version := $(gitsha)
endif
ldflags=-ldflags='-X github.com/mongodb-labs/pcgc/pkg/httpclient.version=$(version)'
build:
	go build $(ldflags) ./...
.PHONY: build

# GIT hooks
link-git-hooks:
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;
.PHONY: link-git-hooks
