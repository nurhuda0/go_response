package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware memvalidasi Bearer token
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Mengambil Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Mengecek apakah token valid
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != "custom_token_1234567890abcdefgHIJKLMNOPQRSTUVWXYZ" { // token yang dianggap valid
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Jika token valid, lanjutkan ke handler berikutnya
		return c.Next()
	}
}
