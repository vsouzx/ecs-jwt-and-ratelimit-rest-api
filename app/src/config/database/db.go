package database

import (
	"fmt"
	"log"
	"os"

	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MustConnect() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name,
	)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("erro ao conectar no MySQL: %v", err)
	}

	if os.Getenv("RUN_AUTOMIGRATE") == "true" {
		if err := db.AutoMigrate(&model.User{}); err != nil {
			log.Fatalf("erro ao migrar modelos: %v", err)
		}
	}

	return db
}
