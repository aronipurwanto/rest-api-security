package handlers

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-security-mod06-logging/utils"
)

// CreateOrder handler untuk menangani pembuatan order
func CreateOrder(c *fiber.Ctx) error {
	payload := c.Body()

	// Simulate error jika payload kosong
	if len(payload) == 0 {
		utils.Log.WithFields(utils.Fields{
			"error": "Payload kosong",
		}).Error("Kesalahan Internal Server")

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"code":   "500",
			"data":   fiber.Map{"message": "Internal Server Error"},
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"code":   "200",
		"data":   fiber.Map{"message": "Order diterima"},
	})
}
