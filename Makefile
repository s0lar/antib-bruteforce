CI_COMMIT_SHA ?= local
CGO_ENABLED = 0
GOARCH = amd64
LDFLAGS = -ldflags "-X main.shaCommit=${CI_COMMIT_SHA}"
GO = $(shell which go)
GO_BUILD = GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GO) build $(LDFLAGS)

.PHONY: echo
echo:
	echo ${CI_COMMIT_SHA}

.PHONY: test
test:
	go test -p 1 -v -race ./...

.PHONY: test-coverage
test-coverage:
	go test -p 1 -v -race -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: test-coverage-open
test-coverage-open: test-coverage
	open cover.html

.PHONY: lint
lint:
	revive ./...

.PHONY: protoc
protoc:
	mkdir -p ./gen
	protoc --proto_path=. --go_out=plugins=grpc:./gen ./proto/service.proto

.PHONY: build-server
build-server:
	$(GO_BUILD) -o ./bin/server ./cmd/server

.PHONY: build-mars-consumer
build-mars-consumer:
	$(GO_BUILD) -o ./bin/mars_consumer ./cmd/mars_consumer

.PHONY: build-wtis-consumer
build-wtis-consumer:
	$(GO_BUILD) -o ./bin/wtis_consumer ./cmd/wtis_consumer
