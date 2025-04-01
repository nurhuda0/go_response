package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware is a middleware for validating Bearer token
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Check if the token is valid
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != "your_static_token" { // Ganti dengan token yang valid
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Token valid, lanjutkan ke handler berikutnya
		return c.Next()
	}
}