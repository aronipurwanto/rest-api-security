package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// Profile handler untuk endpoint /profile
func Profile(c *fiber.Ctx) error {
	log.Println("Mengakses /profile")
	return c.JSON(fiber.Map{
		"message": "Welcome to your profile!",
	})
}

// Settings handler untuk endpoint /settings
func Settings(c *fiber.Ctx) error {
	log.Println("Mengakses /settings")
	return c.JSON(fiber.Map{
		"message": "Here are your settings.",
	})
}

// Dashboard handler untuk endpoint /dashboard
func Dashboard(c *fiber.Ctx) error {
	log.Println("Mengakses /dashboard")
	return c.JSON(fiber.Map{
		"message": "Welcome to your dashboard!",
	})
}
