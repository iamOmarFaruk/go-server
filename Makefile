.PHONY: build run dev test clean docker-build docker-run docker-stop help

# Application name
APP_NAME := go-web-app
BINARY := main
PORT := 8080

# Go build settings
GOOS := linux
GOARCH := amd64
CGO_ENABLED := 0

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Local development targets
deps: ## Download Go dependencies
	go mod download
	go mod tidy

run: ## Run the application locally
	DEV_MODE=true go run cmd/server/main.go

dev: ## Run the application with hot reload (requires air)
	@if ! command -v air &> /dev/null; then \
		echo "Installing air for hot reload..."; \
		go install github.com/cosmtrek/air@latest; \
	fi
	DEV_MODE=true air

test: ## Run tests
	go test -v ./...

clean: ## Clean build artifacts
	rm -f $(BINARY)
	go clean

build: ## Build the binary for production
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY) ./cmd/server

# Docker targets
docker-build: ## Build Docker image
	docker build -t $(APP_NAME) .

docker-run: docker-build ## Run Docker container
	docker run -p $(PORT):$(PORT) --name $(APP_NAME) $(APP_NAME)

docker-stop: ## Stop and remove Docker container
	docker stop $(APP_NAME) 2>/dev/null || true
	docker rm $(APP_NAME) 2>/dev/null || true

docker-compose-build: ## Build and start with docker-compose
	docker compose up --build

docker-compose-down: ## Stop docker-compose
	docker compose down

docker-compose-dev: ## Start docker-compose in development mode with volume mounts
	docker compose -f docker-compose.dev.yml up --build

# Production targets
build-all: ## Build for multiple platforms
	CGO_ENABLED=0 go build -o $(BINARY)-linux-amd64 -ldflags="-w -s" ./cmd/server
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BINARY)-darwin-amd64 -ldflags="-w -s" ./cmd/server
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BINARY)-windows-amd64.exe -ldflags="-w -s" ./cmd/server

# Utility targets
format: ## Format Go code
	go fmt ./...

lint: ## Run Go linter
	@if command -v golangci-lint &> /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

mod-update: ## Update Go modules
	go get -u ./...
	go mod tidy

# Development server
server: deps run ## Install dependencies and run server