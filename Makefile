VERSION_SRC = src/github.com/zylisp/gisp/gitcommit.go
LAST_TAG = $(shell git describe --abbrev=0 --tags)
LAST_COMMIT = $(shell git rev-parse --short HEAD)
DOC_DIR = doc/doc
GODOC=godoc -index -links=true -notes="BUG|TODO|XXX|ISSUE"

.PHONY: build test all

all: build test build-examples test-cli

deps:
	go get github.com/op/go-logging

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

gogen-examples:
	zylisp -cli -go -dir ./bin/examples examples/*.gsp

ast-examples:
	zylisp -cli -ast -dir ./examples examples/*.gsp

bin/examples/%: bin/examples/%.go
	go build -o $@ $<

build-examples: gogen-examples ast-examples
	@$(MAKE) $(basename $(wildcard ./bin/examples/*.go))
	rm ./bin/examples/*.go

clean-examples:
	rm -rf ./bin/examples
	rm ./examples/*.ast

test-cli:
	./tests/test-zylisp-cli.sh

bench-inner-outer:
	go test -v -run=^$ -bench=. ./play/func_call_benchmark_test.go

docs:
	@echo "Generating HTML files ..."
	@echo
	@mkdir -p $(DOC_DIR)/cmd/zylisp $(DOC_DIR)/core $(DOC_DIR)/generator \
					 $(DOC_DIR)/generator/helpers $(DOC_DIR)/lexer \
					 $(DOC_DIR)/parser $(DOC_DIR)/repl
	@$(GODOC) -url /pkg/github.com/zylisp/gisp > \
		$(DOC_DIR)/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/gisp/cmd/ > \
		$(DOC_DIR)/cmd/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/gisp/cmd/zylisp > \
		$(DOC_DIR)/cmd/zylisp/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/gisp/core > \
		$(DOC_DIR)/core/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/gisp/generator > \
		$(DOC_DIR)/generator/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/gisp/generator/helpers > \
		$(DOC_DIR)/generator/helpers/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/gisp/lexer > \
		$(DOC_DIR)/lexer/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/gisp/parser > \
		$(DOC_DIR)/parser/index.html
	@$(GODOC) -url /pkg/github.com/zylisp/gisp/repl > \
		$(DOC_DIR)/repl/index.html

view-docs: docs
	@echo "View project docs in a browser at:"
	@echo "  http://localhost:6060/pkg/"
	@echo "In particular, the zylisp command docs are here:"
	@echo "  http://localhost:6060/pkg/github.com/zylisp/gisp/cmd/zylisp/"
	@echo
	@echo "Starting docs HTTP server ..."
	@GOPATH=`pwd` $(GODOC) -http=:6060 -goroot=`pwd`/doc

publish-docs: docs
	@cd doc && \
	git commit -am "Regen'ed docs." && \
	git push origin gh-pages
