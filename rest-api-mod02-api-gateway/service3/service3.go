// service3/service3.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

const API_KEY = "service3-key"

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

	// Example route for customers
	app.Get("/customers", func(c *fiber.Ctx) error {
		customers := []fiber.Map{
			{"id": 1, "name": "Customer 1"},
			{"id": 2, "name": "Customer 2"},
		}
		return c.JSON(fiber.Map{
			"code":    "200",
			"message": "Customer list retrieved",
			"data":    customers,
		})
	})

	log.Fatal(app.Listen(":3003"))
}
