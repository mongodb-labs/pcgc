TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build

test: fmtcheck
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
	golint -set_exit_status ./...
	GOGC=30 golangci-lint run ./...

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./package-name"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: build test testacc vet fmt fmtcheck errcheck lint test-compile

clean:
	rm -rf out
	go clean ./...

gitsha := $(shell git log -n1 --pretty='%h')
version=$(shell git describe --exact-match --tags "$(gitsha)" 2>/dev/null)
ifeq ($(version),)
	version := $(gitsha)
endif
build: fmtcheck errcheck lint
	mkdir -p out
	go build -ldflags='-X "github.com/mongodb-labs/pcgc/pkg/httpclient=$(version)"' ./...
