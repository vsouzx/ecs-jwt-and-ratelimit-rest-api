package redis

import (
	"context"
	"os"
	"strconv"

	redis "github.com/redis/go-redis/v9"
)

func MustConnect() *redis.Client {
	endpoint := os.Getenv("REDIS_ENDPOINT")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASS")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     endpoint + ":" + port,
		Password: password,
		DB:       db,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic("erro ao conectar no Redis: " + err.Error())
	}

	return redisClient
}
