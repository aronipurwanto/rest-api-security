package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod02-ip-whitelist/controllers"
	"rest-api-mod02-ip-whitelist/middlewares"
)

func main() {
	// Inisialisasi Fiber app
	app := fiber.New()

	// Daftar IP yang diizinkan (whitelist)
	allowedIPs := []string{
		"127.0.0.1",    // Localhost (IPv4)
		"::1",          // Localhost (IPv6)
		"192.168.1.10", // Contoh IP diizinkan lainnya
	}

	// Apply IP Whitelist middleware
	app.Use(middlewares.IPWhitelistMiddleware(allowedIPs))

	// Inisialisasi ResourceController
	resourceController := controllers.ResourceController{}

	// Mendefinisikan route untuk mengambil resource
	app.Get("/resource", resourceController.GetResourceHandler)

	// Menjalankan server di port 3000
	fmt.Println("Server running on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
