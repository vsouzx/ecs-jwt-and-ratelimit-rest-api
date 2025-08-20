package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/controller"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/middleware"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/repository"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/service/auth"
	"gorm.io/gorm"
)

func BuildServer(db *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{})

	app.Use(middleware.AuthJWT())

	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK")})

	//Injecao de dependecia sem o framework
	userRepository := repository.NewUserRepository(db)
	authService := auth.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	// Rotas públicas
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)

	fmt.Println("Rotas registradas.")
	return app
}
