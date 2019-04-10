GOLANGCI_LINT=$(DEFAULT_GOBIN)/golangci-lint

$(GOLANGCI_LINT):
	@echo "Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v1.15.0

vet:
	GO111MODULE=on go vet $(PROJ)/
	GO111MODULE=on go vet $(PROJ)/cmd/zylisp
	# GO111MODULE=on go vet $(PROJ)/core
	# GO111MODULE=on go vet $(PROJ)/generator
	# GO111MODULE=on go vet $(PROJ)/lexer
	# GO111MODULE=on go vet $(PROJ)/parser
	GO111MODULE=on go vet $(PROJ)/repl

goimports:
	goimports -v -w ./
	
lint-all: $(GOLANGCI_LINT)
	@echo "\nLinting source code ...\n"
	@GO111MODULE=off $(GOLANGCI_LINT) run ./

lint-cmd: $(GOLANGCI_LINT)
	@cd ./cmd/zylisp && \
	GO111MODULE=off $(GOLANGCI_LINT) run

lint-repl: $(GOLANGCI_LINT)
	@cd ./repl && \
	GO111MODULE=off $(GOLANGCI_LINT) run

lint: lint-repl lint-cmd