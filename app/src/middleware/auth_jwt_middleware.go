package middleware

import (
	"strings"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT() fiber.Handler {
	skip := map[string]struct{}{
		"/health": {},
		"/register": {},
		"/login": {},
	}

	return func(c *fiber.Ctx) error {

		// Skip para rotas públicas
		if _, ok := skip[c.Path()]; ok {
			return c.Next()
		}
		auth := c.Get("Authorization")
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return fiber.NewError(fiber.StatusUnauthorized)
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "algoritmo inválido")
			}
			jwtSecret := os.Getenv("JWT_SECRET")
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized)
		}

		// (Opcional) Poderíamos extrair claims e salvar em c.Locals("userID", ...)
		return c.Next()
	}
}
