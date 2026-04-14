# ECS JWT & Rate Limit REST API

A production-ready Go REST API featuring JWT authentication and Redis-backed rate limiting, deployed to AWS ECS Fargate via Terraform and GitHub Actions CI/CD.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-ECS_Fargate-FF9900?style=flat&logo=amazon-aws&logoColor=white)
![Terraform](https://img.shields.io/badge/Terraform-IaC-7B42BC?style=flat&logo=terraform&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=flat&logo=docker&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-Rate_Limiting-DC382D?style=flat&logo=redis&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?style=flat&logo=mysql&logoColor=white)

## Overview

This project demonstrates a cloud-native REST API architecture with:

- **JWT Authentication** — HS256-signed tokens with 15-minute expiration
- **Redis Rate Limiting** — sliding window counter per client IP with fail-open behavior
- **AWS ECS Fargate** — containerized workload with autoscaling (1–4 tasks based on CPU)
- **Infrastructure as Code** — fully provisioned with Terraform, state in S3 + DynamoDB locking
- **CI/CD Pipeline** — GitHub Actions builds and deploys on every push to `main`



### AWS Infrastructure

```
AWS
├── VPC
├── Application Load Balancer
├── ECS Fargate Cluster
│   └── Service (desired: 2, min: 1, max: 4 — autoscales at CPU 50%)
├── RDS MySQL 8.0 (db.t3.micro)
├── ElastiCache Redis 7.1 (cache.t3.micro)
└── CloudWatch Logs
```

### Rate Limit Headers

Every response includes:

```
X-RateLimit-Limit: 10
X-RateLimit-Remaining: 7
```

When the limit is exceeded, the API returns `429 Too Many Requests`.

## Local Development

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- MySQL and Redis running locally

### Running with Docker Compose

```bash
cd app
docker compose up
```

### Running directly

```bash
cd app

# Set environment variables
source setup-local.sh

# Run the server
go run ./cmd/api
```

## Project Structure

```
.
├── app/
│   ├── cmd/api/          # Entry point
│   ├── internal/
│   │   ├── config/       # DB and Redis initialization
│   │   ├── handler/      # HTTP handlers
│   │   ├── middleware/   # JWT auth + rate limiter
│   │   ├── model/        # GORM models
│   │   ├── repository/   # Database queries
│   │   ├── service/      # Business logic
│   │   └── server/       # Fiber app setup and routing
│   ├── Dockerfile
│   └── docker-compose.yml
└── infra/                # Terraform (VPC, ALB, ECS, RDS, ElastiCache)
```

State is stored remotely in S3 with DynamoDB locking.

## CI/CD

Pushing to `main` triggers the GitHub Actions pipeline:

1. Builds and pushes a multi-platform Docker image (`linux/arm64`) to Docker Hub
2. Runs `terraform init`, selects the `prod` workspace, and applies infrastructure changes
