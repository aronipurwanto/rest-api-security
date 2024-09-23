package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

const apiKey = "your-api-key-12345" // Gantilah ini dengan API Key yang valid

// Middleware untuk validasi API Key
func apiKeyMiddleware(c *fiber.Ctx) error {
	clientAPIKey := c.Get("X-API-Key")

	// Cek apakah API Key sesuai
	if clientAPIKey != apiKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid API Key",
		})
	}

	// Lanjutkan jika API Key valid
	return c.Next()
}

func main() {
	app := fiber.New()

	// Route public tanpa API Key
	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is a public endpoint, no API Key required.",
		})
	})

	// Route protected dengan API Key middleware
	app.Use("/protected", apiKeyMiddleware)

	// Endpoint yang dilindungi
	app.Get("/protected/data", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "This is a protected endpoint, API Key is valid.",
		})
	})

	// Jalankan server di port 3000
	log.Fatal(app.Listen(":3000"))
}
