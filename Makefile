VERSION=0.9.0-dev
PROJ=github.com/zylisp/zylisp
PACKAGE=$(PROJ)/common
BUILD_FLAGS=$(shell govvv -flags -pkg $(PACKAGE) -version $(VERSION))
BIN=bin
ZY=./$(BIN)/zylisp
ZYC=./$(BIN)/zyc
DEFAULT_GOPATH=$(shell tr ':' '\n' <<< $$GOPATH|awk '!x[$$0]++'|sed '/^$$/d'|head -1)
ifeq ($(DEFAULT_GOPATH),)
DEFAULT_GOPATH := ~/go
endif
DEFAULT_GOBIN=$(DEFAULT_GOPATH)/bin
export PATH:=$(PATH):$(DEFAULT_GOBIN)

.PHONY: build test all

default: all

include internal/make/docker.mk
include internal/make/docs.mk
include internal/make/lint.mk
include internal/make/mods.mk
include internal/make/test.mk

all: clean build lint-all test build-examples clean-examples test-cli \
		test-examples clean-examples test-zyc

default-gopath:
	@echo $(DEFAULT_GOPATH)

deps:
	@GO111MODULE=on go get github.com/sirupsen/logrus@v1.4.0
	@GO111MODULE=on go install github.com/sirupsen/logrus

build: deps
	@echo "\nBuilding zylisp ...\n"
	@GO111MODULE=on \
	go build -ldflags="$(BUILD_FLAGS)" $(PROJ)
	@GO111MODULE=on GOBIN=`pwd`/$(BIN) \
	go install -ldflags="$(BUILD_FLAGS)" $(PROJ)/cmd/zylisp

install-zyc:
	@mkdir -p ~/go/bin
	@cp bin/zyc ~/go/bin

install-zylisp:
	@go get $(PROJ)/cmd/zylisp

install: install-zylisp install-zyc

clean-go:
	@echo "Removing Go object files, etc. ..."
	@go clean

clean-bin:
	@echo "Removing untracked files from ./$(BIN) ..."
	@git ls-files $(BIN) --others | xargs rm -f

clean: clean-go clean-bin clean-examples

clean-all: clean clean-modules
