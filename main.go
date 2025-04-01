package main

import (
	"go_response/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // Menginisiasi aplikasi dengan Fiber

	// Endpoint secara publik
	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Ini adalah endpoint publik.",
		})
	})

	// Endpoint dengan middleware
	app.Get("/protected", middleware.AuthMiddleware(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Ini adalah endpoint yang dilindungi.",
		})
	})

	// Menjalankan aplikasi
	app.Listen(":3000")
}
