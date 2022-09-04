CI_COMMIT_SHA ?= local
CGO_ENABLED = 0
GOARCH = amd64
LDFLAGS = -ldflags "-X main.shaCommit=${CI_COMMIT_SHA}"
GO = $(shell which go)
GO_BUILD = GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GO) build $(LDFLAGS)

.PHONY: test
test: ## Run tests
	go test -v -race -count 100 ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	go test -p 1 -v -race -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: docker-up
docker-up: ## Run test containers
	docker-compose up -d

.PHONY: docker-down
docker-down: ## Stop test containers
	docker-compose down

.PHONY: lint
lint: ## Run linter
	CGO_ENABLED=0 golangci-lint run

.PHONY: protoc
protoc: ## Generate protoc
	mkdir -p ./gen
	protoc --proto_path=. --go_out=./gen --go-grpc_out=./gen ./proto/service.proto

.PHONY: build
build:  ## Build
	$(GO_BUILD) -o ./bin/server ./cmd/server
	$(GO_BUILD) -o ./bin/cli ./cmd/cli

.PHONY: build-server
build-server:  ## Build server
	$(GO_BUILD) -o ./bin/server ./cmd/server
	$(GO_BUILD) -o ./bin/cli ./cmd/cli
	./bin/server

.PHONY: run
run:  ## Run server
	docker-compose up -d
	$(GO_BUILD) -o ./bin/server ./cmd/server && ./bin/server

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'