VERSION = 1.0.0-alpha3
BUILD_FLAGS=$(shell govvv -flags -pkg github.com/zylisp/zylisp -version $(VERSION))
DOC_DIR = doc/doc
GODOC=godoc -index -links=true -notes="BUG|TODO|XXX|ISSUE"

.PHONY: build test all

all: build lint-all test build-examples test-cli test-examples test-zyc clean-examples

travis: lint-deps test-deps all

deps:
	go get github.com/ahmetb/govvv
	go get github.com/zylisp/zylog/logger
	go get github.com/sirupsen/logrus

lint-deps:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b ~/go/bin v1.15.0
	go get golang.org/x/tools/cmd/goimports

test-deps:
	go get github.com/masukomi/check

build: deps
	@go install -ldflags="$(BUILD_FLAGS)" ./cmd/zylisp

lint-all:
	golangci-lint run ./

lint-cmd:
	cd .p/cmd/zylisp && \
	golangci-lint run

lint-repl:
	cd ./repl && \
	golangci-lint run

lint: lint-repl lint-cmd

vet:
	go vet github.com/zylisp/zylisp/
	go vet github.com/zylisp/zylisp/cmd/zylisp
	# go vet github.com/zylisp/zylisp/core
	# go vet github.com/zylisp/zylisp/generator
	# go vet github.com/zylisp/zylisp/lexer
	# go vet github.com/zylisp/zylisp/parser
	go vet github.com/zylisp/zylisp/repl

test: test-deps
	@echo "running 'go test' for core ..." && \
	cd ./core && \
	go test -v && \
	cd - && \
	echo "running 'go test' for lexer ..." && \
	cd ./lexer && \
	go test -v
	@$(MAKE) clean-examples

gogen-examples:
	zylisp -cli -go -dir ./bin/examples examples/*.zsp

ast-examples:
	zylisp -cli -ast -dir ./examples examples/*.zsp

bin/examples/%: bin/examples/%.go
	go build -o $@ $<

build-examples: gogen-examples ast-examples
	@$(MAKE) $(basename $(wildcard ./bin/examples/*.go))
	rm ./bin/examples/*.go

clean-examples:
	rm -rf ./bin/examples
	rm -f ./examples/*.ast
	rm -f ./examples/*.go

test-cli:
	./tests/test-zylisp-cli.sh

test-examples:
	./tests/test-compiled-examples.sh

test-zyc:
	./tests/test-zyc.sh

bench-inner-outer:
	go test -v -run=^$ -bench=. ./play/func_call_benchmark_test.go

docs:
	@echo "Generating HTML files ..."
	@echo
	@mkdir -p $(DOC_DIR)/cmd/zylisp $(DOC_DIR)/core $(DOC_DIR)/generator \
					 $(DOC_DIR)/generator/helpers $(DOC_DIR)/lexer \
					 $(DOC_DIR)/parser $(DOC_DIR)/repl $(DOC_DIR)/common \
					 $(DOC_DIR)/util
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp > \
		$(DOC_DIR)/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/cmd/ > \
		$(DOC_DIR)/cmd/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/cmd/zylisp > \
		$(DOC_DIR)/cmd/zylisp/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/core > \
		$(DOC_DIR)/core/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/generator > \
		$(DOC_DIR)/generator/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/generator/helpers > \
		$(DOC_DIR)/generator/helpers/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/lexer > \
		$(DOC_DIR)/lexer/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/parser > \
		$(DOC_DIR)/parser/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/repl > \
		$(DOC_DIR)/repl/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/zylisp/common > \
		$(DOC_DIR)/common/index.html

view-docs: docs
	@echo "View project docs in a browser at:"
	@echo "  http://localhost:6060/pkg/"
	@echo "In particular, the zylisp command docs are here:"
	@echo "  http://localhost:6060/pkg/github.com/zylisp/zylisp/cmd/zylisp/"
	@echo
	@echo "Starting docs HTTP server ..."
	@GOPATH=`pwd` $(GODOC) -http=:6060 -goroot=`pwd`/doc

publish-docs: docs
	@cd doc && \
	git commit -am "Regen'ed docs." && \
	git push origin gh-pages

install-zyc:
	@mkdir -p ~/go/bin
	@cp bin/zyc ~/go/bin

install-zylisp:
	@go get github.com/zylisp/zylisp/cmd/zylisp

install: install-zylisp install-zyc

goimports:
	GO111MODULE=on goimports -v -w ./
