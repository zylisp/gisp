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