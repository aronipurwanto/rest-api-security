package handlers

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod03-encoding/utils"
)

// AesEncryptDecrypt encrypts and decrypts a string using AES from request
func AesEncryptDecrypt(c *fiber.Ctx) error {
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

	key := []byte("a very very very very secret key") // 32 bytes key for AES-256
	log.Println("AES encryption process started for data:", req.Data)

	// Encrypt the data
	encryptedData, err := utils.EncryptAES(key, req.Data)
	if err != nil {
		log.Fatal("Error during AES encryption:", err)
	}

	// Convert encrypted data to base64
	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedData)
	log.Println("AES encryption completed. Encrypted data:", encryptedBase64)

	// Decrypt the data
	decryptedData, err := utils.DecryptAES(key, encryptedData)
	if err != nil {
		log.Fatal("Error during AES decryption:", err)
	}

	log.Println("AES decryption completed. Decrypted data:", decryptedData)

	// Return formatted response
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "AES encryption and decryption successful",
		"data": fiber.Map{
			"encrypted": encryptedBase64,
			"decrypted": decryptedData,
		},
	})
}
