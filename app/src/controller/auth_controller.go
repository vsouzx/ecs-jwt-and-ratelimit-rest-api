package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/dto"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/service/auth"
)

type AuthController struct {
	authService auth.AuthServiceInterface
}

func NewAuthController(authService auth.AuthServiceInterface) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// POST /register (público)
func (ac *AuthController) Register(c *fiber.Ctx) error {
	var request dto.RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Payload invalido")
	}
	return ac.authService.Register(c, request)
}

// POST /login (público)
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var request dto.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Payload invalido")
	}
	return ac.authService.Login(c, request)
}
