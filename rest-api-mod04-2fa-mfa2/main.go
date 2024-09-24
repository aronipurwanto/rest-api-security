package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod04-2fa-mfa2/middlewares"
	"rest-api-mod04-2fa-mfa2/routes"
)

func main() {
	app := fiber.New()

	// Gunakan middleware logger untuk mencatat setiap permintaan
	app.Use(middlewares.RequestLogger())

	// Setup routes
	routes.SetupRoutes(app)

	// Jalankan server di port 3000
	log.Println("Server berjalan di http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
