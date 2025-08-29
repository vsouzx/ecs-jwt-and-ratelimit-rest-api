resource "aws_lb" "load_balancer" {
  name               = "${var.app_name}-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb.id]
  subnets            = data.aws_subnets.default_in_vpc.ids
}

resource "aws_lb_target_group" "target_group" {
  name        = "${var.app_name}-tg"
  port        = var.container_port
  protocol    = "HTTP"
  vpc_id      = data.aws_vpc.default.id
  target_type = "ip" # obrigatório para Fargate

  health_check {
    enabled  = true
    path     = "/health" # ajuste se necessário
    matcher  = "200-399"
    interval = 30
    timeout  = 5
  }
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.load_balancer.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.target_group.arn
  }
}

