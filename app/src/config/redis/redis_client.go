package redis

import (
	"context"
	"os"
	"strconv"

	redis "github.com/redis/go-redis/v9"
)

func MustConnect() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASS")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic("erro ao conectar no Redis: " + err.Error())
	}

	return redisClient
}
