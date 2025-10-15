# ========================================
# 💼 AIpply Makefile
# ========================================

APP_NAME := AIpply
BINARY := bin/$(APP_NAME)
PKG := ./...
GOFILES := $(shell find . -name '*.go' -not -path "./vendor/*")

# Default target
.DEFAULT_GOAL := help

# Environment
ENV ?= dev
GIN_MODE ?= debug

# Colors for logs
YELLOW := \033[1;33m
GREEN := \033[0;32m
RED := \033[0;31m
RESET := \033[0m

help:
	@echo "$(YELLOW)Available commands:$(RESET)"
	@echo "  make dev             🔁 Run dev server with Air (hot reload)"
	@echo "  make lint            🧹 Run Revive linter"
	@echo "  make test            🧪 Run Go tests"
	@echo "  make tidy            🧩 Run go mod tidy"
	@echo "  make build           🏗️  Build binary"
	@echo "  make run             🚀 Run binary in release mode"
	@echo "  make clean           🧼 Remove build artifacts"

# Run app in dev mode using Air (requires .air.toml)
dev:
	@echo "$(GREEN)Starting dev server with Air...$(RESET)"
	@air || echo "$(RED)❌ Air not installed. Run: go install github.com/cosmtrek/air@latest$(RESET)"

# Run revive linter (requires revive.toml config)
lint:
	@echo "$(GREEN)Linting code with Revive...$(RESET)"
	@revive -config revive.toml ./... || echo "$(RED)Linting failed!$(RESET)"

# Run tests
test:
	@echo "$(GREEN)Running tests...$(RESET)"
	@go test -v $(PKG)

# Tidy up dependencies
tidy:
	@echo "$(GREEN)Tidying up modules...$(RESET)"
	@go mod tidy

# Build production binary
build:
	@echo "$(GREEN)Building $(APP_NAME)...$(RESET)"
	@go build -o $(BINARY) .

# Run in production mode
run:
	@echo "$(GREEN)Running $(APP_NAME) in release mode...$(RESET)"
	@ENV=prod GIN_MODE=release ./$(BINARY)

# Clean build artifacts
clean:
	@echo "$(GREEN)Cleaning build artifacts...$(RESET)"
	@rm -rf $(BINARY) tmp
