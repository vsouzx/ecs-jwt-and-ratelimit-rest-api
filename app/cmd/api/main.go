package main

import (
	"log"

	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/internal/config"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/internal/server"
)

func main() {
	db := config.MustConnectDB()
	redisClient := config.MustConnectRedis()

	app := server.Build(db, redisClient)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
