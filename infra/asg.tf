resource "aws_appautoscaling_target" "asg-gofiber-service" {
  max_capacity       = 4
  min_capacity       = 1
  resource_id        = "service/${aws_ecs_cluster.gofiber-api-cluster.name}/${aws_ecs_service.gofiber-api.name}"
  scalable_dimension = "ecs:service:DesiredCount"
  service_namespace  = "ecs"
}

resource "aws_appautoscaling_policy" "cpu_target" {
  name               = "${var.app_name}-cpu-target"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.asg-gofiber-service.resource_id
  scalable_dimension = aws_appautoscaling_target.asg-gofiber-service.scalable_dimension
  service_namespace  = aws_appautoscaling_target.asg-gofiber-service.service_namespace


  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageCPUUtilization"
    }
    target_value       = 50
    scale_in_cooldown  = 60
    scale_out_cooldown = 60
  }
}
