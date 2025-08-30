resource "aws_elasticache_subnet_group" "redis_subnet_group" {
  name       = "redis-default-vpc-subnet-group"
  subnet_ids = data.aws_subnets.default_in_vpc.ids
}

resource "aws_elasticache_replication_group" "redis" {
  replication_group_id       = "my-redis-rg"
  description                = "Redis com AUTH habilitado"
  engine                     = "redis"
  engine_version             = "7.1"
  node_type                  = "cache.t3.micro"
  port                       = 6379

  num_node_groups            = 1            # 1 shard
  replicas_per_node_group    = 0            # sem réplicas

  transit_encryption_enabled = true
  at_rest_encryption_enabled = true
  auth_token                 = "admin123@"

  automatic_failover_enabled = false        # sem réplicas => precisa ser false

  subnet_group_name          = aws_elasticache_subnet_group.redis_subnet_group.name
  security_group_ids         = [aws_security_group.redis_sg.id]
  parameter_group_name       = "default.redis7"
}
