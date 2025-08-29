package main

import (
	"log"

	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/config/database"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/config/redis"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/router"
)

func main() {
	db := database.MustConnect()
	redisClient := redis.MustConnect()

	app := router.BuildServer(db, redisClient)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
