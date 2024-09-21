package handler

import "github.com/gofiber/fiber/v2"

// SecureDataHandler handles requests to the /secure-data endpoint
func SecureDataHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Access granted to secure data!"})
}
