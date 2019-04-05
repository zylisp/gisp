VERSION=1.0.0-alpha4
PROJ=github.com/zylisp/zylisp
PACKAGE=$(PROJ)/common
BUILD_FLAGS=$(shell govvv -flags -pkg $(PACKAGE) -version $(VERSION))
BIN=bin
ZY=./$(BIN)/zylisp
ZYC=./$(BIN)/zyc
DOC_DIR=doc/doc
GODOC=godoc -index -links=true -notes="BUG|TODO|XXX|ISSUE"
GOLANGCI_LINT=$(shell which golangci-lint)
DEFAULT_GOPATH=$(shell tr ':' '\n' <<< "$$GOPATH"|sed '/^$$/d'|head -1)
DEFAULT_GOBIN=$(DEFAULT_GOPATH)/bin
DOCKER_ORG=zylisp
DOCKER_TAG=zylisp
DOCKER_BINARY=$(ZY)-linux

.PHONY: build test all

default: all

all: clean build lint-all test build-examples clean-examples test-cli test-examples clean-examples test-zyc

default-gopath:
	@echo $(DEFAULT_GOPATH)

deps:
	@GO111MODULE=on go get github.com/sirupsen/logrus@v1.4.0
	@GO111MODULE=on go install github.com/sirupsen/logrus

$(GOLANGCI_LINT):
	@echo "Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v1.15.0

build: deps
	@echo "\nBuilding zylisp ...\n"
	@GO111MODULE=on \
	go build -ldflags="$(BUILD_FLAGS)" $(PROJ)
	@GO111MODULE=on GOBIN=`pwd`/$(BIN) \
	go install -ldflags="$(BUILD_FLAGS)" $(PROJ)/cmd/zylisp

docker-binary:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
	go build -ldflags="$(BUILD_FLAGS)" -o $(DOCKER_BINARY) \
	$(PROJ)/cmd/zylisp

lint-all: $(GOLANGCI_LINT)
	@echo "\nLinting source code ...\n"
	@golangci-lint run ./

lint-cmd: $(GOLANGCI_LINT)
	cd .p/cmd/zylisp && \
	golangci-lint run

lint-repl: $(GOLANGCI_LINT)
	cd ./repl && \
	golangci-lint run

lint: lint-repl lint-cmd

vet:
	GO111MODULE=on go vet $(PROJ)/
	GO111MODULE=on go vet $(PROJ)/cmd/zylisp
	# GO111MODULE=on go vet $(PROJ)/core
	# GO111MODULE=on go vet $(PROJ)/generator
	# GO111MODULE=on go vet $(PROJ)/lexer
	# GO111MODULE=on go vet $(PROJ)/parser
	GO111MODULE=on go vet $(PROJ)/repl

test:
	@echo "\nRunning 'go test' for core ...\n" && \
	cd ./core && \
	GO111MODULE=on go test -v
	@echo "\nRunning 'go test' for lexer ...\n" && \
	cd ./lexer && \
	GO111MODULE=on go test -v

gogen-examples:
	@echo "\nGenerating .go files for examples ...\n"
	$(ZY) -cli -go -dir ./bin/examples examples/*.zsp

ast-examples:
	@echo "\nGenerating AST files for examples ...\n"
	$(ZY) -cli -ast -dir ./examples examples/*.zsp

bin/examples/%: bin/examples/%.go
	GO111MODULE=on go build -o $@ $<

build-examples: gogen-examples ast-examples
	@echo "\nCompiling example go files ...\n"
	@$(MAKE) $(basename $(wildcard ./bin/examples/*.go))

clean-examples:
	@echo "Removing generated example files ..."
	@rm -rf ./bin/examples ./examples/*.ast ./examples/*.go

test-cli:
	./tests/test-zylisp-cli.sh

test-examples:
	./tests/test-compiled-examples.sh

test-zyc:
	./tests/test-zyc.sh

bench-inner-outer:
	go test -v -run=^$ -bench=. ./play/func_call_benchmark_test.go

docs:
	@echo "\nGenerating HTML files ...\n"
	@mkdir -p $(DOC_DIR)/cmd/zylisp $(DOC_DIR)/core $(DOC_DIR)/generator \
					 $(DOC_DIR)/generator/helpers $(DOC_DIR)/lexer \
					 $(DOC_DIR)/parser $(DOC_DIR)/repl $(DOC_DIR)/common \
					 $(DOC_DIR)/util
	@$(GODOC) -url /pkg/$(PROJ) > \
		$(DOC_DIR)/index.html
	@$(GODOC) -url /pkg/$(PROJ)/cmd/ > \
		$(DOC_DIR)/cmd/index.html
	@$(GODOC) -url /pkg/$(PROJ)/cmd/zylisp > \
		$(DOC_DIR)/cmd/zylisp/index.html
	@$(GODOC) -url /pkg/$(PROJ)/core > \
		$(DOC_DIR)/core/index.html
	@$(GODOC) -url /pkg/$(PROJ)/generator > \
		$(DOC_DIR)/generator/index.html
	@$(GODOC) -url /pkg/$(PROJ)/generator/helpers > \
		$(DOC_DIR)/generator/helpers/index.html
	@$(GODOC) -url /pkg/$(PROJ)/lexer > \
		$(DOC_DIR)/lexer/index.html
	@$(GODOC) -url /pkg/$(PROJ)/parser > \
		$(DOC_DIR)/parser/index.html
	@$(GODOC) -url /pkg/$(PROJ)/repl > \
		$(DOC_DIR)/repl/index.html
	@$(GODOC) -url /pkg/$(PROJ)/common > \
		$(DOC_DIR)/common/index.html

view-docs: docs
	@echo "\nView project docs in a browser at:"
	@echo "  http://localhost:6060/pkg/"
	@echo "In particular, the zylisp command docs are here:"
	@echo "  http://localhost:6060/pkg/$(PROJ)/cmd/zylisp/\n"
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
	@go get $(PROJ)/cmd/zylisp

install: install-zylisp install-zyc

goimports:
	goimports -v -w ./

list-packages:
	GO111MODULE=on go list ./...

modules-init:
	GO111MODULE=on go mod init $(PROJ)/cmd/zylisp

modules-update:
	GO111MODULE=on go get -u

module-upgrades:
	GO111MODULE=on go list -u -m all

modules-tidy:
	GO111MODULE=on go mod tidy

clean-modules:
	GO111MODULE=on go clean --modcache

clean-go:
	@echo "Removing Go object files, etc. ..."
	@go clean

clean-bin:
	@echo "Removing untracked files from ./$(BIN) ..."
	@git ls-files $(BIN) --others | xargs rm -f

clean: clean-go clean-bin clean-examples

clean-all: clean clean-modules

docker-img: docker-binary
	docker build -t $(DOCKER_ORG)/$(DOCKER_TAG):$(VERSION) .
	rm -rf $(DOCKER_BINARY)

run-img:
	docker run \
		-it $(DOCKER_ORG)/$(DOCKER_TAG):$(VERSION)
