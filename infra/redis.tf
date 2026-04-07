resource "aws_elasticache_subnet_group" "main" {
  name       = "${var.app_name}-redis-subnet-group"
  subnet_ids = data.aws_subnets.default_in_vpc.ids
}

resource "aws_elasticache_replication_group" "main" {
  replication_group_id       = "${var.app_name}-redis"
  description                = "Redis with AUTH enabled"
  engine                     = "redis"
  engine_version             = "7.1"
  node_type                  = "cache.t3.micro"
  port                       = 6379

  num_node_groups         = 1 # 1 shard
  replicas_per_node_group = 0 # no replicas

  transit_encryption_enabled = true
  at_rest_encryption_enabled = true
  auth_token                 = local.app_secrets.redis_auth_token

  automatic_failover_enabled = false # requires at least 2 nodes

  subnet_group_name  = aws_elasticache_subnet_group.main.name
  security_group_ids = [aws_security_group.redis.id]
  parameter_group_name = "default.redis7"
}
