variable "aws_region" {
  description = "AWS region where all resources will be provisioned"
  type        = string
  default     = "us-east-1"
}

variable "app_name" {
  description = "Base name used as prefix for all resource names"
  type        = string
  default     = "gofiber-api"
}

variable "image_uri" {
  description = "Full container image URI (Docker Hub, ECR, etc.) injected by CI/CD via TF_VAR_image_uri"
  type        = string
  default     = "vsouzx/ecs-jwt-and-ratelimit-rest-api:latest"
}

variable "cpu_architecture" {
  description = "CPU architecture for the ECS task runtime platform (must match the container image architecture)"
  type        = string
  default     = "ARM64"

  validation {
    condition     = contains(["ARM64", "X86_64"], var.cpu_architecture)
    error_message = "cpu_architecture must be ARM64 or X86_64."
  }
}

variable "container_port" {
  description = "Port exposed by the container"
  type        = number
  default     = 8080
}

variable "db_secret_arn" {
  description = "ARN of the AWS Secrets Manager secret containing db_username and db_password"
  type        = string
}

variable "db_name" {
  description = "Name of the MySQL database to create"
  type        = string
  default     = "app"
}

variable "app_secrets_arn" {
  description = "ARN of the AWS Secrets Manager secret containing jwt_secret and redis_auth_token"
  type        = string
}

variable "redis_db" {
  description = "Redis database index used by the application"
  type        = number
  default     = 0
}

variable "rate_limit_count" {
  description = "Number of requests allowed in the rate limiting window"
  type        = number
  default     = 10
}

variable "rate_limit_ttl" {
  description = "Rate limiting window duration in minutes"
  type        = number
  default     = 1
}

variable "run_automigrate" {
  description = "Whether the API should run database auto-migrations on startup"
  type        = bool
  default     = false
}

variable "extra_container_env" {
  description = "Additional environment variables to inject into the container"
  type        = map(string)
  default     = {}
}
