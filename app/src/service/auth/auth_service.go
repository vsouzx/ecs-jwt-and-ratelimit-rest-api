package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/dto"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/model"
	"github.com/vsouzx/ecs-jwt-ratelimit-rest-api/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(ctx *fiber.Ctx, req dto.RegisterRequest) error
	Login(ctx *fiber.Ctx, req dto.LoginRequest) error
}

type authService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewAuthService(userRepository repository.UserRepositoryInterface) *authService {
	return &authService{
		UserRepository: userRepository,
	}
}

func (as *authService) Register(ctx *fiber.Ctx, req dto.RegisterRequest) error {
	existingUser, err := as.UserRepository.FindByEmail(req.Email)
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
	if err := as.UserRepository.Create(newUser); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Erro salvar usuário no banco: %s", err.Error()))
	}

	print("a")

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Usuário criado com sucesso"})
}

func (as *authService) Login(ctx *fiber.Ctx, req dto.LoginRequest) error {
	user, err := as.UserRepository.FindByEmail(req.Email)
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
		"exp":        time.Now().Add(time.Duration(15) * time.Minute).Unix(),
		"iat":        time.Now().Unix(),
	}

	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Erro ao gerar token: %s", err.Error())})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
