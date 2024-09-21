package handlers

import "github.com/gofiber/fiber/v2"

// Handler untuk Admin endpoint (CRUD access)
func AdminHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Admin access granted"})
}

// Handler untuk Editor endpoint (READ/UPDATE access)
func EditorHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Editor access granted"})
}

// Handler untuk Viewer endpoint (READ access)
func ViewerHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Viewer access granted"})
}
