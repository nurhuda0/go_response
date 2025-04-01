package main

import (
	"github.com/gofiber/fiber/v2"
	"go_response/middleware"
)

func main() {
	app := fiber.New()

	// Endpoint publik
	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is a public endpoint",
		})
	})

	// Endpoint dengan middleware
	app.Get("/protected", middleware.AuthMiddleware(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is a protected endpoint",
		})
	})

	// Jalankan aplikasi
	app.Listen(":3000")
}