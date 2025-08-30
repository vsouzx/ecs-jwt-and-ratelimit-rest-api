resource "aws_iam_role" "ecs_task_execution" {
  name = "${var.app_name}-task-exec-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_attach" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

########################################
# ECS Cluster + Task Definition + Service
########################################

resource "aws_ecs_cluster" "gofiber-api-cluster" {
  name = "${var.app_name}-cluster"
}

locals {
  container_env = [for k, v in var.container_env : {
    name  = k
    value = v
  }]
}

resource "aws_ecs_task_definition" "task_definition" {
  family                   = "${var.app_name}-taskdef"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256" # 0.25 vCPU
  memory                   = "512" # 0.5 GB
  execution_role_arn       = aws_iam_role.ecs_task_execution.arn

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "ARM64" # ou "ARM64" se você realmente quiser ARM
  }

  container_definitions = jsonencode([
    {
      name      = var.app_name
      image     = var.image
      essential = true
      portMappings = [
        {
          containerPort = var.container_port
          hostPort      = var.container_port
          protocol      = "tcp"
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = aws_cloudwatch_log_group.this.name
          awslogs-region        = var.aws_region
          awslogs-stream-prefix = "ecs"
        }
      }
      environment = local.container_env
    }
  ])
}

resource "aws_ecs_service" "gofiber-api" {
  name            = "${var.app_name}-service"
  cluster         = aws_ecs_cluster.gofiber-api-cluster.arn
  task_definition = aws_ecs_task_definition.task_definition.arn
  desired_count   = 2
  launch_type     = "FARGATE"


  network_configuration {
    subnets          = data.aws_subnets.default_in_vpc.ids
    security_groups  = [aws_security_group.task.id]
    assign_public_ip = true # usa IP público para sair à internet
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.target_group.arn
    container_name   = var.app_name
    container_port   = var.container_port
  }

  depends_on = [aws_lb_listener.http]
}
