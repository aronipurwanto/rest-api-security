// service2/service2.go
package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

const API_KEY = "service2-key"

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

	// Example route for orders
	app.Get("/orders", func(c *fiber.Ctx) error {
		orders := []fiber.Map{
			{"id": 1, "product": "Product 1", "quantity": 2},
			{"id": 2, "product": "Product 2", "quantity": 1},
		}
		return c.JSON(fiber.Map{
			"code":    "200",
			"message": "Order list retrieved",
			"data":    orders,
		})
	})

	log.Fatal(app.Listen(":3002"))
}
