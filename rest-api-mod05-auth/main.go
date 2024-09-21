package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"log"
	"rest-api-mod05-auth/config"
	"rest-api-mod05-auth/handlers"
	"rest-api-mod05-auth/middleware"
	"rest-api-mod05-auth/models"
)

func main() {
	// Memuat konfigurasi menggunakan Viper
	config.LoadConfig()

	// Hubungkan ke database MySQL
	config.ConnectDB()

	// Migrasi model ke database
	config.DB.AutoMigrate(&models.User{}, &models.Role{})

	// Inisialisasi data
	models.SeedData(config.DB)

	// Inisialisasi Fiber
	app := fiber.New()

	// Gunakan logger untuk mencatat request
	app.Use(logger.New())

	// Middleware untuk autentikasi
	app.Use(middleware.AuthMiddleware)

	// Route untuk admin (hanya admin dengan full CRUD access)
	app.Get("/admin", middleware.RoleMiddleware("admin"), handlers.AdminHandler)

	// Route untuk editor (hanya editor dengan UPDATE/READ access)
	app.Get("/editor", middleware.AuthorityMiddleware("UPDATE"), handlers.EditorHandler)

	// Route untuk viewer (hanya viewer dengan READ access)
	app.Get("/viewer", middleware.AuthorityMiddleware("READ"), handlers.ViewerHandler)

	// Jalankan server di port yang ditentukan dalam konfigurasi
	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	log.Fatal(app.Listen(port))
}
