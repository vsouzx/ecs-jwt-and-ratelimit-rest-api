package service

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/internal/model"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthService interface {
	Register(ctx *fiber.Ctx, req RegisterRequest) error
	Login(ctx *fiber.Ctx, req LoginRequest) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(ctx *fiber.Ctx, req RegisterRequest) error {
	existingUser, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Erro ao buscar usuário pelo e-mail %s: %s", req.Email, err.Error()))
	}

	if existingUser != (model.User{}) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Já existe um usuário com este e-mail"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Erro ao criptografar senha: %s", err.Error()))
	}

	newUser := &model.User{Name: req.Name, Password: string(hashedPassword), Email: req.Email}
	if err := s.userRepo.Create(newUser); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Erro salvar usuário no banco: %s", err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Usuário criado com sucesso"})
}

func (s *authService) Login(ctx *fiber.Ctx, req LoginRequest) error {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Erro ao buscar usuário pelo e-mail %s: %s", req.Email, err.Error()))
	}

	if user == (model.User{}) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Credenciais inválidas."})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Credenciais inválidas."})
	}

	claims := jwt.MapClaims{
		"identifier": user.Identifier,
		"exp":        time.Now().Add(15 * time.Minute).Unix(),
		"iat":        time.Now().Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Erro ao gerar token: %s", err.Error())})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
