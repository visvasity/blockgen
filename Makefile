# Copyright (c) 2025 Visvasity LLC

export GO ?= $(shell which go)
export GOBIN = $(CURDIR)
export PATH = $(CURDIR):$(HOME)/bin:$(HOME)/go/bin:/bin:/usr/bin:/usr/local/bin:/sbin:/usr/sbin
export GOTESTFLAGS ?=

.PHONY: all
all:
	$(GO) generate ./...
	# goimports -w ./output/*.go || true
	$(GO) build ./...
	$(GO)	test -fullpath -failfast -count=1 -coverprofile=coverage.out $(GOTESTFLAGS) -v ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

.PHONY: clean
clean:
	git clean -f -X
