package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-security-mod06-logging/handlers"
	"rest-api-security-mod06-logging/middleware"
)

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Gunakan middleware logging
	app.Use(middleware.LoggingMiddleware)

	// Routes
	app.Post("/api/v1/orders", handlers.CreateOrder)
	app.Post("/api/v1/login", handlers.Login)
	app.Get("/api/v1/admin", handlers.AccessAdmin)

	// Start server
	log.Println("Server berjalan di port 3000")
	log.Fatal(app.Listen(":3000"))
}
