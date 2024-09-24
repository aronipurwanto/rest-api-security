package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod04-basic-auth/handlers"
	"rest-api-mod04-basic-auth/middleware"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Logging to indicate server start
	log.Println("Starting server on port 3000...")

	// Routes with Basic Authentication middleware applied
	app.Post("/login", middleware.BasicAuth, handlers.LoginHandler)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
