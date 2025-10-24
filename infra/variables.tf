variable "aws_region" {
    type = string
    default = "us-east-1"
}

variable "app_name" {
    description = "Nome base dos recursos"
    type = string
    default = "gofiber-api"
}

variable "image" {
    description = "Imagem do container (Docker Hub, ECR, etc.)"
    type = string
    default = "vsouzx/ecs-jwt-and-ratelimit-rest-api:latest"
}

variable "container_port" {
    description = "Porta exposta pelo container"
    type = number
    default = 8080
}

variable "db_secret_arn" {
  type = string
}

variable "jwt_secret" {
  description = "JWT secret used by the application"
  type        = string
  sensitive   = true
}

variable "redis_auth_token" {
  description = "Authentication token for the Redis replication group"
  type        = string
  sensitive   = true
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
  description = "Rate limiting window in seconds"
  type        = number
  default     = 1
}

variable "run_automigrate" {
  description = "Whether the API should run database automigrations on startup"
  type        = bool
  default     = false
}

variable "extra_container_env" {
  description = "Additional environment variables to inject into the container"
  type        = map(string)
  default     = {}
}
