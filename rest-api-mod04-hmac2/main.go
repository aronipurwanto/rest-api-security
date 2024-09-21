package main

import (
	"fmt"
	"github.com/aronipurwanto/rest-api-mod04-hmac2/handler"
	"github.com/aronipurwanto/rest-api-mod04-hmac2/middleware"
	"github.com/aronipurwanto/rest-api-mod04-hmac2/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"time"
)

func main() {
	// Start a Fiber app
	app := fiber.New()

	// Use logger middleware
	app.Use(logger.New())

	// Apply HMAC middleware for secure routes
	app.Use(middleware.HMACMiddleware)

	// Secure route
	app.Get("/secure-data", handler.SecureDataHandler)

	// Simulate a client request (in a real case, this will be done by a client application)
	simulateClientRequest()

	// Start the server
	log.Fatal(app.Listen(":3000"))
}

// Simulate a client generating the HMAC and timestamp for the request
func simulateClientRequest() {
	// Simulate a timestamp
	timestamp := time.Now().Format(time.RFC3339)

	// Generate HMAC from the timestamp
	hmacSignature := service.GenerateHMAC(timestamp, service.SecretKey)

	// Output the required headers for the request
	fmt.Println("X-HMAC-Signature:", hmacSignature)
	fmt.Println("X-Timestamp:", timestamp)
}
