# SG do Load Balancer (HTTP 80 público)
resource "aws_security_group" "alb" {
    name = "${var.app_name}-alb-sg"
    description = "ALB SG"
    vpc_id = data.aws_vpc.default.id

    ingress {
        description = "HTTP"
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
        ipv6_cidr_blocks = ["::/0"]
    }


    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
        ipv6_cidr_blocks = ["::/0"]
    }
}

# SG das Tasks (apenas tráfego do ALB na porta do container)
resource "aws_security_group" "task" {
    name = "${var.app_name}-task-sg"
    description = "ECS Task SG"
    vpc_id = data.aws_vpc.default.id

    ingress {
        description = "ALB to Task"
        from_port = var.container_port
        to_port = var.container_port
        protocol = "tcp"
        security_groups = [aws_security_group.alb.id]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
        ipv6_cidr_blocks = ["::/0"]
    }   
}

resource "aws_security_group" "rds_sg" {
  name        = "rds-mysql-sg"
  description = "Allow MySQL access"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description = "Allow MySQL from my IP"
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    #security_groups = [aws_security_group.lambda_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "rds-mysql-sg"
  }
}


resource "aws_security_group" "redis_sg" {
  name        = "redis-sg"
  description = "Acesso ao Redis apenas a partir das ECS tasks"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description     = "Redis 6379 das ECS tasks"
    from_port       = 6379
    to_port         = 6379
    protocol        = "tcp"
    security_groups = [aws_security_group.task.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}