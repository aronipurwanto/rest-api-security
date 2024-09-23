package handlers

import (
	"crypto/sha256"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Sha256Hash hashes a string using SHA256 from request
func Sha256Hash(c *fiber.Ctx) error {
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

	log.Println("SHA256 hashing process started for data:", req.Data)

	// Create a new SHA256 hash
	hash := sha256.New()
	hash.Write([]byte(req.Data))

	// Get the hashed result
	hashedData := hash.Sum(nil)

	// Convert hashed data to hexadecimal string
	hashedString := fmt.Sprintf("%x", hashedData)

	log.Println("SHA256 hashing completed. Hashed data:", hashedString)

	// Return formatted response
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "SHA256 hashing successful",
		"data":    fiber.Map{"hashed": hashedString},
	})
}
