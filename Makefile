# Makefile for Go project with cross-compilation and common commands

# Project parameters
BINARY_NAME = kanban
SOURCE_DIR = ./
SOURCE_FILE = $(SOURCE_DIR)/cmd/server/main.go

# Go build parameters
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED = 0

# Cross-compilation targets
TARGETS = darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64

.PHONY: all clean build run test vet lint fmt deps cross-compile help

all: help

help:  ## Show this help message
	@echo "Usage: make [TARGET]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Variables:"
	@printf "  \033[36mGOOS\033[0m            The OS to build for (default: current OS)\n"
	@printf "  \033[36mGOARCH\033[0m        The architecture to build for (default: current architecture)\n"
	@echo ""

clean: ## Remove the binary and other build artifacts
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe
	@echo "Done."

build: ## Build the binary for the current OS and architecture
	@echo "Building..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME) $(SOURCE_FILE)
	@echo "Build complete. The executable is $(BINARY_NAME)"

run: ## Run the application
	@echo "Running..."
	go run $(SOURCE_FILE)

test: ## Run tests
	@echo "Testing..."
	go test ./...

vet: ## Run go vet
	@echo "Vetting..."
	go vet ./...

lint: ## Run golint (ensure golint is installed: go install golang.org/x/lint/golint@latest)
	@echo "Linting..."
	golint ./...

fmt: ## Run gofmt
	@echo "Formatting..."
	gofmt -w .

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download

cross-compile: ## Cross-compile for multiple targets
	@for target in $(TARGETS); do \
		GOOS=$${target%%/*} GOARCH=$${target##*/} \
		echo "Cross-compiling for $${target%%/*}/$${target##*/}..."; \
		CGO_ENABLED=$(CGO_ENABLED) GOOS=$${target%%/*} GOARCH=$${target##*/} \
		go build -o $(BINARY_NAME)-$${target%%/*}-$${target##*/} $(SOURCE_FILE); \
	done
	@echo "Cross-compilation complete."
