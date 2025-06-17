# Microservices Assignment

This repository contains two microservices built using GoCommons libraries and patterns:

## Services

### 1. Order Management Service (OMS)
- CSV order processing via S3 and SQS
- Kafka event streaming for order lifecycle
- MongoDB for order storage
- REST APIs with authentication and logging
- Webhook registration and updates

### 2. Inventory Management Service (IMS)
- Hub and SKU management
- Inventory operations with atomic updates
- PostgreSQL for data storage
- Redis caching for validation
- REST APIs for CRUD operations

## Tech Stack

- **GoCommons**: HTTP server, Kafka, Redis, SQS, S3, PostgreSQL, logging, config
- **MongoDB**: Order storage (standard Go driver)
- **Docker**: Containerization and local development
- **Docker Compose**: Local infrastructure setup

## Development Setup

1. Set GOPRIVATE environment variable:
   ```bash
   go env -w GOPRIVATE="github.com/omniful/"
   ```

2. Configure Git for SSH access to private repositories

3. Run local infrastructure:
   ```bash
   docker-compose up -d
   ```

4. Start the services:
   ```bash
   # Start IMS first (OMS depends on it)
   go run cmd/ims/main.go
   
   # Start OMS
   go run cmd/oms/main.go
   ```

## Project Structure

```
assignment/
├── cmd/
│   ├── oms/                    # OMS main application
│   └── ims/                    # IMS main application
├── internal/
│   ├── oms/                    # OMS business logic
│   ├── ims/                    # IMS business logic
│   └── common/                 # Shared utilities
├── configs/                    # Configuration files
├── docker/                     # Docker configurations
└── docs/                       # Documentation
```

## Status

🚧 **In Development** - Step by step implementation in progress 