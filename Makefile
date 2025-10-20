.PHONY: help dev db-up db-down db-migrate sqlc build run clean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

dev: ## Start development environment
	@echo "Starting development environment..."
	@make db-up
	@echo "Waiting for database to be ready..."
	@sleep 3
	@make db-migrate
	@echo "Development environment ready!"

db-up: ## Start PostgreSQL database
	docker-compose up -d

db-down: ## Stop PostgreSQL database
	docker-compose down

db-migrate: ## Run database migrations
	@echo "Running database migrations..."
	@PGPASSWORD=pony psql -h localhost -U pony -d pony -f db/schema.sql

sqlc: ## Generate sqlc code
	sqlc generate

build: ## Build the application
	go build -o bin/pony ./cmd/pony

run: ## Run the application
	go run ./cmd/pony/main.go

clean: ## Clean build artifacts
	rm -rf bin/
	go clean

install-tools: ## Install required tools
	@echo "Installing sqlc..."
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@echo "Tools installed!"
