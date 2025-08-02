.PHONY: help install build test run clean docker-build docker-run

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	@echo "Installing Go dependencies..."
	cd backend && go mod tidy
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

build: ## Build all services
	@echo "Building backend services..."
	cd backend && go build -o bin/user-service ./cmd/user-service
	cd backend && go build -o bin/expense-service ./cmd/expense-service
	cd backend && go build -o bin/investment-service ./cmd/investment-service
	cd backend && go build -o bin/goal-service ./cmd/goal-service
	cd backend && go build -o bin/report-service ./cmd/report-service
	cd backend && go build -o bin/api-gateway ./cmd/api-gateway
	@echo "Building frontend..."
	cd frontend && npm run build

test: ## Run tests
	@echo "Running backend tests..."
	cd backend && go test ./...
	@echo "Running frontend tests..."
	cd frontend && npm test

run: ## Run all services locally
	@echo "Starting services..."
	docker-compose up -d

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -rf backend/bin/*
	rm -rf frontend/build/*
	rm -rf frontend/node_modules

docker-build: ## Build Docker images
	@echo "Building Docker images..."
	docker-compose build

docker-run: ## Run with Docker
	@echo "Starting with Docker..."
	docker-compose up -d

docker-stop: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	docker-compose down

logs: ## View logs
	docker-compose logs -f

migrate: ## Run database migrations
	@echo "Running database migrations..."
	cd backend && go run cmd/migrate/main.go 