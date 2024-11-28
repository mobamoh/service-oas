# ==============================================================================
# Modules support
tidy:
	go mod tidy
	go mod vendor

fmt:
	go fmt ./...

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

# ==============================================================================
# Class Stuff

run:
	go run api/services/sales/main.go | go run api/tooling/logfmt/main.go

generate: ## generate code with ogen based on openapi specs
	go generate ./...

# ==============================================================================
# Building containers

GOM_CORE_APP    := service-oas
BASE_IMAGE_NAME := localhost/service-oas
VERSION         := 0.0.1
GOM_CORE_IMAGE  := $(BASE_IMAGE_NAME)/$(GOM_CORE_APP):$(VERSION)

build: service-oas

service-oas:
	docker build \
		-f Dockerfile \
		-t $(GOM_CORE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

# ==============================================================================

# Docker Compose

compose-up:
	docker compose -f docker_compose.yaml -p compose up -d

compose-build-up: build compose-up

compose-down:
	docker compose -f docker_compose.yaml down

compose-logs:
	docker compose -f docker_compose.yaml logs


# ==============================================================================
# Metrics and Tracing

metrics-view-sc:
	expvarmon -ports="localhost:3010" -vars="build,requests,goroutines,errors,panics,mem:memstats.HeapAlloc,mem:memstats.HeapSys,mem:memstats.Sys"

metrics-view:
	expvarmon -ports="localhost:4020" -endpoint="/metrics" -vars="build,requests,goroutines,errors,panics,mem:memstats.HeapAlloc,mem:memstats.HeapSys,mem:memstats.Sys"

grafana:
	open http://localhost:3100/

statsviz:
	open http://localhost:3010/debug/statsviz

# ==============================================================================
# Install dependencies

dev-gotooling:
	go install github.com/ogen-go/ogen/cmd/ogen@latest
	go install github.com/divan/expvarmon@latest
	go install github.com/rakyll/hey@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

dev-brew:
	brew update
	brew list go || brew install go

# ==============================================================================
# Help command
help:
	@echo "Usage: make <command>"
	@echo ""
	@echo "Commands:"
	@echo "  dev-gotooling           Install Go tooling"
	@echo "  dev-brew                Install brew dependencies"
	@echo "  build                   Build all the containers"