package handlers

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-security-mod06-logging/utils"
)

// AccessAdmin handler untuk akses admin
func AccessAdmin(c *fiber.Ctx) error {
	utils.Log.WithFields(utils.Fields{
		"user":   "johndoe",
		"status": 403,
	}).Error("Upaya Akses Tidak Sah")

	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"status": "error",
		"code":   "403",
		"data":   fiber.Map{"message": "Forbidden"},
	})
}
