package config

import (
	"context"
	"crypto/tls"
	"os"
	"strconv"

	redis "github.com/redis/go-redis/v9"
)

func MustConnectRedis() *redis.Client {
	endpoint := os.Getenv("REDIS_ENDPOINT")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASS")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	opts := &redis.Options{
		Addr:      endpoint + ":" + port,
		Password:  password,
		DB:        db,
		TLSConfig: &tls.Config{ServerName: endpoint},
	}

	redisClient := redis.NewClient(opts)

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic("erro ao conectar no Redis: " + err.Error())
	}

	return redisClient
}
