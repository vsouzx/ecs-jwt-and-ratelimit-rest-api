package main

import (
	"log"

	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/config/database"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/router"
)

func main() {
	db := database.MustConnect()

	app := router.BuildServer(db)

	addr := ":3000"
	log.Printf("🚀 Server rodando em %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
