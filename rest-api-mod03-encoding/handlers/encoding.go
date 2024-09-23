package handlers

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Base64Encode encodes a given string into base64 from request
func Base64Encode(c *fiber.Ctx) error {
	// Parse request body
	type Request struct {
		Data string `json:"data"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid request body",
			"data":    nil,
		})
	}

	log.Println("Base64 encoding process started for data:", req.Data)

	// Base64 encoding
	encoded := base64.StdEncoding.EncodeToString([]byte(req.Data))

	log.Println("Base64 encoding completed. Encoded data:", encoded)

	// Return formatted response
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Base64 encoding successful",
		"data":    fiber.Map{"encoded": encoded},
	})
}
