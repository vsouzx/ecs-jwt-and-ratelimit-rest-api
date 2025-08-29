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

variable "container_env" {
    description = "Variáveis de ambiente do container (mapa nome->valor)"
    type = map(string)
    default = {}
}
