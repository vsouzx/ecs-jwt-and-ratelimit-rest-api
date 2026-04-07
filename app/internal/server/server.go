package server

import (
	"github.com/gofiber/fiber/v2"
	redis "github.com/redis/go-redis/v9"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/internal/handler"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/internal/middleware"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/internal/repository"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/internal/service"
	"gorm.io/gorm"
)

func Build(db *gorm.DB, redisClient *redis.Client) *fiber.App {
	app := fiber.New()

	app.Use(middleware.RateLimiter(redisClient))
	app.Use(middleware.AuthJWT())

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	return app
}
