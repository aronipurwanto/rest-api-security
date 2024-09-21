package handlers

import "github.com/gofiber/fiber/v2"

// Handler untuk Admin endpoint (RBAC-based)
func AdminHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Admin access granted"})
}

// Handler untuk Editor endpoint (ACL-based)
func EditorHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Editor access granted"})
}

// Handler untuk Viewer endpoint (ACL-based)
func ViewerHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Viewer access granted"})
}
