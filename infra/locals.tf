locals {
  db_secret = jsondecode(data.aws_secretsmanager_secret_version.db_creds.secret_string)
}

data "aws_secretsmanager_secret_version" "db_creds" {
  secret_id = var.db_secret_arn
}