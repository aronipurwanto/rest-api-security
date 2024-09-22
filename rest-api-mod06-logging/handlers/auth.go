package handlers

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-security-mod06-logging/utils"
)

// Login handler untuk autentikasi pengguna
func Login(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token != "Bearer valid-token" {
		utils.Log.WithFields(utils.Fields{
			"user":  "johndoe",
			"ip":    c.IP(),
			"error": "Token kedaluwarsa",
		}).Warn("Autentikasi Gagal")

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error",
			"code":   "401",
			"data":   fiber.Map{"message": "Unauthorized"},
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"code":   "200",
		"data":   fiber.Map{"message": "Login Berhasil"},
	})
}
