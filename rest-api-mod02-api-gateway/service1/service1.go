// service1/service1.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

const API_KEY = "service1-key"

func main() {
	app := fiber.New()

	// Middleware to check API Key
	app.Use(func(c *fiber.Ctx) error {
		apiKey := c.Get("API-Key")
		if apiKey != API_KEY {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    "401",
				"message": "Invalid API Key",
				"data":    nil,
			})
		}
		return c.Next()
	})

	// Example route for products
	app.Get("/products", func(c *fiber.Ctx) error {
		products := []fiber.Map{
			{"id": 1, "name": "Product 1"},
			{"id": 2, "name": "Product 2"},
		}
		return c.JSON(fiber.Map{
			"code":    "200",
			"message": "Product list retrieved",
			"data":    products,
		})
	})

	log.Fatal(app.Listen(":3001"))
}
