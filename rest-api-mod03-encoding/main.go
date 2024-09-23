package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod03-encoding/handlers"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Logging to indicate that the server is starting
	log.Println("Starting server on port 3000...")

	// Routes
	app.Post("/encode/base64", handlers.Base64Encode)
	app.Post("/hash/sha256", handlers.Sha256Hash)
	app.Post("/encrypt/aes", handlers.AesEncryptDecrypt)

	// Start server on port 3000
	log.Fatal(app.Listen(":3000"))
}
