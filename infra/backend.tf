terraform {
  # Bucket, key, region e dynamodb_table são passados via -backend-config no CI/CD.
  # Terraform não suporta interpolação de variáveis no bloco backend.
  backend "s3" {}
}
