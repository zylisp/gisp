VERSION_SRC = src/github.com/zylisp/gisp/gitcommit.go
LAST_TAG = $(shell git describe --abbrev=0 --tags)
LAST_COMMIT = $(shell git rev-parse --short HEAD)

.PHONY: build test all

all: build test build-examples

deps:

lint-deps:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b ~/go/bin v1.15.0

test-deps:
	go get github.com/masukomi/check

build: deps
	@echo "package gisp" > $(VERSION_SRC)
	@echo "" >> $(VERSION_SRC)
	@echo "func init() { GITLASTTAG = \"$(LAST_TAG)\"; \
		GITLASTCOMMIT = \"$(LAST_COMMIT)\" }" >> $(VERSION_SRC)
	@go install github.com/zylisp/gisp/cmd/zylisp

lint-all: lint-deps
	golangci-lint run

lint-cmd:
	cd src/github.com/zylisp/gisp/cmd/zylisp && \
	golangci-lint run

lint-repl:
	cd src/github.com/zylisp/gisp/repl && \
	golangci-lint run

lint: lint-repl lint-cmd

vet:
	go vet github.com/zylisp/gisp/
	go vet github.com/zylisp/gisp/cmd/zylisp
	# go vet github.com/zylisp/gisp/core
	# go vet github.com/zylisp/gisp/generator
	# go vet github.com/zylisp/gisp/lexer
	# go vet github.com/zylisp/gisp/parser
	go vet github.com/zylisp/gisp/repl


test: test-deps
	echo "running 'go test' for core ..." && \
	cd src/github.com/zylisp/gisp/core && \
	go test -v && \
	cd - && \
	echo "running 'go test' for lexer ..." && \
	cd src/github.com/zylisp/gisp/lexer && \
	go test -v

build-examples:
	zyc -o ./bin/examples examples/*.gsp

clean-examples:
	rm -rf ./bin/examples

bench-inner-outer:
	go test -v -run=^$ -bench=. ./play/func_call_benchmark_test.go
