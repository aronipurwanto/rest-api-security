package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// LoginHandler handles the login logic after successful Basic Authentication
func LoginHandler(c *fiber.Ctx) error {
	log.Println("LoginHandler reached, user authenticated.")

	// Simulating fetching user data after authentication
	user := c.Locals("user").(string)

	// Log the authenticated user
	log.Printf("User %s successfully authenticated", user)

	// Response after successful authentication
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Authentication successful",
		"data":    fiber.Map{"user": user},
	})
}
