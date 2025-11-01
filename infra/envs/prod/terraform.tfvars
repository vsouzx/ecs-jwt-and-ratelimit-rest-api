container_env = {
  DB_USER          = aws_db_instance.default.username
  DB_PASSWORD      = local.db_secret.db_password
  DB_HOST          = aws_db_instance.default.address
  DB_PORT          = tostring(aws_db_instance.default.port)
  DB_NAME          = aws_db_instance.default.db_name
  JWT_SECRET       = "E-zzg4VU#pOHu=qr5rEp&PEBq86Ak#q4UQswhRnV-L6zFsdfSsFnXD-86uQ6maFdAOUpjXzAzznc#KR63mQrB&SDR2QwEB5kRjAhvHhkR3PB88NKZpHh8#D3=eufOjR99KoqxNvhfPst=QPkxR1LoenuuEsE4tQDL70fQEDhv81Kxlh4hpBvpPM7lgxIyEEejVP9=iBbcBnPS=Cll=xAumnh&kOiD3oZ=wM6Qq2xPC#6rJUf5nqUmyYBzakPTpK#27eezrXy6EMOnJOJbau4uwZ9ohuCuWGwmJrx-d6O8snAo3Wq3wlmxwIBAxPkBMkNZJq81f#U0ZerHw-JQxUYWcSq"
  REDIS_ENDPOINT   = aws_elasticache_replication_group.redis.primary_endpoint_address
  REDIS_PORT       = tostring(aws_elasticache_replication_group.redis.port)
  REDIS_PASS       = "admin123@"
  REDIS_DB         = "0"
  RATE_LIMIT_COUNT = "10"
  RATE_LIMIT_TTL   = "1"
}

db_secret_arn="arn:aws:secretsmanager:us-east-1:337328321041:secret:prod-mysql-credentials-Bbqjg5"