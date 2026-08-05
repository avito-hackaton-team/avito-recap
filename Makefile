SHELL := /bin/bash
.DEFAULT_GOAL := help

BACKEND_DIR := backend/recap
BACKEND_GO_MOD := $(BACKEND_DIR)/go.mod
BIN_DIR := $(CURDIR)/bin

GOLANGCI_LINT_VERSION := v2.12.2
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint
GOLANGCI_CONFIG := $(CURDIR)/.golangci.yml
GOOSE_VERSION := v3.27.1
GOOSE := $(BIN_DIR)/goose
OGEN_VERSION := v1.23.0
OGEN := $(BIN_DIR)/ogen

ENV_FILE := $(CURDIR)/.env
MIGRATIONS_DIR := $(CURDIR)/backend/recap/migrations/migrations
OPENAPI_SPEC := $(CURDIR)/backend/recap/api/recap/v1/openapi.yaml
GENERATED_API_DIR := $(CURDIR)/backend/recap/generated/recapapi

.PHONY: help tools require-backend require-env lint-config format lint vet test test-race \
	test-integration tidy tidy-check generate generate-api check up down logs \
	migrate-up migrate-down migrate-status

help:
	@echo "Available commands:"
	@echo "  make tools         Install development tools"
	@echo "  make format        Format Go code"
	@echo "  make lint          Run golangci-lint"
	@echo "  make vet           Run go vet"
	@echo "  make test          Run unit tests"
	@echo "  make test-race     Run unit tests with the race detector"
	@echo "  make test-integration Run integration tests when present"
	@echo "  make tidy          Synchronize Go dependencies"
	@echo "  make tidy-check    Check whether go.mod and go.sum are tidy"
	@echo "  make generate      Generate code from project contracts"
	@echo "  make check         Run all required Go checks"
	@echo "  make up            Start local infrastructure"
	@echo "  make down          Stop local infrastructure"
	@echo "  make logs          Follow infrastructure logs"
	@echo "  make migrate-up    Apply database migrations"
	@echo "  make migrate-down  Roll back the latest migration"
	@echo "  make migrate-status Show database migration status"

$(GOLANGCI_LINT):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install \
		github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

$(GOOSE):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install \
		github.com/pressly/goose/v3/cmd/goose@$(GOOSE_VERSION)

$(OGEN):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install \
		github.com/ogen-go/ogen/cmd/ogen@$(OGEN_VERSION)

tools: $(GOLANGCI_LINT) $(GOOSE) $(OGEN)

require-backend:
	@test -f $(BACKEND_GO_MOD) || { \
		echo "Backend module not found: $(BACKEND_GO_MOD)"; \
		echo "Create it before running Go checks."; \
		exit 1; \
	}

require-env:
	@test -f $(ENV_FILE) || { \
		echo "Environment file not found: $(ENV_FILE)"; \
		echo "Run: cp .env.example .env"; \
		exit 1; \
	}
	@set -a; source $(ENV_FILE); set +a; \
	for variable in POSTGRES_USER POSTGRES_PASSWORD POSTGRES_DB; do \
		if [ -z "$${!variable}" ]; then \
			echo "Required environment variable is empty: $$variable"; \
			exit 1; \
		fi; \
	done

lint-config: require-backend $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) config verify --config=$(GOLANGCI_CONFIG)

format: require-backend $(GOLANGCI_LINT)
	cd $(BACKEND_DIR) && $(GOLANGCI_LINT) fmt --config=$(GOLANGCI_CONFIG)

lint: require-backend $(GOLANGCI_LINT)
	cd $(BACKEND_DIR) && $(GOLANGCI_LINT) run --config=$(GOLANGCI_CONFIG) ./...

vet: require-backend
	cd $(BACKEND_DIR) && go vet ./...

test: require-backend
	cd $(BACKEND_DIR) && go test ./...

test-race: require-backend
	cd $(BACKEND_DIR) && go test -race -count=1 ./...

test-integration: require-backend
	@set -euo pipefail; \
	tests="$$(find $(BACKEND_DIR) -type f -name '*_test.go' \
		-exec grep -l '^//go:build integration' {} + 2>/dev/null || true)"; \
	if [ -z "$$tests" ]; then \
		echo "No integration tests found; skipping."; \
		exit 0; \
	fi; \
	cd $(BACKEND_DIR) && go test -race -count=1 -tags=integration ./...

tidy: require-backend
	cd $(BACKEND_DIR) && go mod tidy

tidy-check: require-backend
	cd $(BACKEND_DIR) && go mod tidy -diff

generate: generate-api

generate-api: require-backend $(OGEN)
	$(OGEN) --target $(GENERATED_API_DIR) --package recapapi --clean $(OPENAPI_SPEC)

check: lint-config lint vet tidy-check test-race test-integration

up: require-env
	docker compose up -d --wait

down:
	docker compose down

logs:
	docker compose logs -f

migrate-up: require-env $(GOOSE)
	@set -a; source $(ENV_FILE); set +a; \
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres \
		"host=$${POSTGRES_BIND_HOST:-127.0.0.1} port=$${POSTGRES_EXTERNAL_PORT:-5432} user=$${POSTGRES_USER} password=$${POSTGRES_PASSWORD} dbname=$${POSTGRES_DB} sslmode=$${POSTGRES_SSL_MODE:-disable}" \
		up

migrate-down: require-env $(GOOSE)
	@set -a; source $(ENV_FILE); set +a; \
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres \
		"host=$${POSTGRES_BIND_HOST:-127.0.0.1} port=$${POSTGRES_EXTERNAL_PORT:-5432} user=$${POSTGRES_USER} password=$${POSTGRES_PASSWORD} dbname=$${POSTGRES_DB} sslmode=$${POSTGRES_SSL_MODE:-disable}" \
		down

migrate-status: require-env $(GOOSE)
	@set -a; source $(ENV_FILE); set +a; \
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres \
		"host=$${POSTGRES_BIND_HOST:-127.0.0.1} port=$${POSTGRES_EXTERNAL_PORT:-5432} user=$${POSTGRES_USER} password=$${POSTGRES_PASSWORD} dbname=$${POSTGRES_DB} sslmode=$${POSTGRES_SSL_MODE:-disable}" \
		status
