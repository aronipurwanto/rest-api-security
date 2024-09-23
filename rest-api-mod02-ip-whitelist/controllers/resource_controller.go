package controllers

import "github.com/gofiber/fiber/v2"

// ResourceController handles resource-related routes
type ResourceController struct{}

// GetResourceHandler handles retrieving a resource
func (rc *ResourceController) GetResourceHandler(c *fiber.Ctx) error {
	// Simulasi pengambilan resource
	resource := map[string]interface{}{
		"id":   1,
		"name": "Sample Resource",
	}

	// Mengembalikan respons dengan format standar
	return c.JSON(fiber.Map{
		"code":    "200",
		"message": "Resource retrieved successfully",
		"data":    resource,
	})
}
