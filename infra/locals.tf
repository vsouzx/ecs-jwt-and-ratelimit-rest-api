locals {
  db_secret   = jsondecode(data.aws_secretsmanager_secret_version.db_creds.secret_string)
  app_secrets = jsondecode(data.aws_secretsmanager_secret_version.app_secrets.secret_string)
}

data "aws_secretsmanager_secret_version" "db_creds" {
  secret_id = var.db_secret_arn
}

# Esperado: { "jwt_secret": "...", "redis_auth_token": "..." }
data "aws_secretsmanager_secret_version" "app_secrets" {
  secret_id = var.app_secrets_arn
}
