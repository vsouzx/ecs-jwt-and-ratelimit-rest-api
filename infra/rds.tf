resource "aws_db_instance" "main" {
  identifier             = "${var.app_name}-db"
  allocated_storage      = 10
  db_name                = var.db_name
  engine                 = "mysql"
  engine_version         = "8.0.41"
  instance_class         = "db.t3.micro"
  username               = local.db_secret.db_username
  password               = local.db_secret.db_password
  parameter_group_name   = "default.mysql8.0"
  skip_final_snapshot    = true
  backup_retention_period = 7
  vpc_security_group_ids = [aws_security_group.rds.id]
  publicly_accessible    = false
}
