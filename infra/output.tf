output "alb_dns_name" {
  description = "DNS do Load Balancer"
  value       = aws_lb.load_balancer.dns_name
}

output "cluster_name" {
  value = aws_ecs_cluster.gofiber-api-cluster.name
}

output "service_name" {
  value = aws_ecs_service.gofiber-api.name
}

output "task_definition_arn" {
  value = aws_ecs_task_definition.task_definition.arn
}
