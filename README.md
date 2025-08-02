# Personal Finance Management System

A comprehensive microservices-based personal finance application built with Go and modern web technologies.

## Architecture

- **Backend**: Go microservices using Go-kit
- **Frontend**: React with TypeScript and Tailwind CSS
- **Database**: PostgreSQL
- **Containerization**: Docker & Docker Compose
- **API**: RESTful APIs

## Services

1. **User Service** - Authentication and user management
2. **Expense Service** - Daily expense tracking and categorization
3. **Investment Service** - Investment portfolio management
4. **Goal Service** - Financial goal setting and tracking
5. **Report Service** - Analytics and reporting
6. **API Gateway** - Centralized API management

## Features

- Daily expense tracking with categorization
- Investment portfolio management (FDs, RDs, Stocks, Mutual Funds, US Stocks)
- Financial goal setting and tracking
- Monthly/Yearly expense reports
- Interactive dashboard with analytics
- Multi-user support

## Quick Start

```bash
# Clone the repository
git clone <your-repo-url>
cd tgfinance

# Start all services
docker-compose up -d

# Access the application
# Frontend: http://localhost:3000
# API Gateway: http://localhost:8080
```

## Development Setup

```bash
# Install dependencies
make install

# Run tests
make test

# Build all services
make build

# Run locally
make run
```

## Project Structure

```
tgfinance/
├── backend/
│   ├── cmd/                    # Application entry points
│   ├── internal/               # Private application code
│   ├── pkg/                    # Public libraries
│   └── services/               # Microservices
├── frontend/                   # React application
├── docker/                     # Docker configurations
├── scripts/                    # Build and deployment scripts
├── docs/                       # Documentation
└── deploy/                     # Deployment configurations
``` 