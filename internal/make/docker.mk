DOCKER_ORG=zylisp
DOCKER_TAG=zylisp
DOCKER_BINARY=$(ZY)-linux

docker-binary:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
	go build -ldflags="$(BUILD_FLAGS)" -o $(DOCKER_BINARY) \
	$(PROJ)/cmd/zylisp

docker-img: docker-binary
	docker build -t $(DOCKER_ORG)/$(DOCKER_TAG):$(VERSION) .
	docker build -t $(DOCKER_ORG)/$(DOCKER_TAG):latest .
	rm -rf $(DOCKER_BINARY)

run-img-ast:
	@docker run \
		-it $(DOCKER_ORG)/$(DOCKER_TAG):$(VERSION) -ast

publish-img: docker-img
	docker push $(DOCKER_ORG)/$(DOCKER_TAG):$(VERSION)
	docker push $(DOCKER_ORG)/$(DOCKER_TAG):latest

clean-docker:
	@docker system prune