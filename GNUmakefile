TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build

test: fmtcheck
	@echo "==> Running tests..."
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

lint:
	@echo "==> Linting all packages..."
	golint -set_exit_status ./...
	GOGC=30 golangci-lint run ./...

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./package-name"; \
		exit 1; \
	fi
	@echo "==> Compiling test binary..."
	go test -c $(TEST) $(TESTARGS)

.PHONY: build install test testacc vet fmt fmtcheck errcheck lint test-compile init-git-hooks

# Build targets
clean:
	@echo "==> Cleaning build artifacts..."
	rm -rf out
	go clean ./...

gitsha := $(shell git log -n1 --pretty='%h')
version=$(shell git describe --exact-match --tags "$(gitsha)" 2>/dev/null)
ifeq ($(version),)
	version := $(gitsha)
endif
ldflags=-ldflags='-X github.com/mongodb-labs/pcgc/pkg/httpclient.version=$(version) -X github.com/mongodb-labs/pcgc/pkg/mpc/cmd.version=$(version)'
build: fmtcheck errcheck lint test
	@echo "==> Building binary for the current architecture..."
	mkdir -p out
	go build $(ldflags) -o out/pcgc

install: fmtcheck errcheck lint test
	@echo "==> Installing pcgc in $(GOPATH)/bin ..."
	go install $(ldflags) ./pkg/mpc

# GIT hooks
link-git-hooks:
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;
