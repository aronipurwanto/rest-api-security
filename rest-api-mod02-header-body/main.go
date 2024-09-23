package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"rest-api-mod02-header-body/controllers"
)

func main() {
	// Inisialisasi aplikasi Fiber
	app := fiber.New()

	// Inisialisasi HeaderBodyController
	headerBodyController := controllers.HeaderBodyController{}

	// Mendefinisikan routing untuk berbagai endpoint
	app.Post("/resource", headerBodyController.CreateResourceHandler)       // Endpoint untuk membuat resource
	app.Get("/resource/:id", headerBodyController.GetResourceHandler)       // Endpoint untuk mengambil resource berdasarkan ID
	app.Put("/resource/:id", headerBodyController.UpdateResourceHandler)    // Endpoint untuk memperbarui resource berdasarkan ID
	app.Delete("/resource/:id", headerBodyController.DeleteResourceHandler) // Endpoint untuk menghapus resource berdasarkan ID
	app.Get("/resource/search", headerBodyController.SearchResourceHandler) // Endpoint untuk mencari resource

	// Menjalankan server di port 3000
	fmt.Println("Server running on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
