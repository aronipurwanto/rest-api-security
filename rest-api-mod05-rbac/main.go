package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"rest-api-mod05-rbac/config"
	"rest-api-mod05-rbac/handlers"
	"rest-api-mod05-rbac/middleware"
	"rest-api-mod05-rbac/models"
)

func main() {
	// Memuat konfigurasi menggunakan Viper
	config.LoadConfig()

	// Hubungkan ke database MySQL
	config.ConnectDB()

	// Migrasi model ke database
	config.DB.AutoMigrate(&models.User{}, &models.Role{}, &models.AccessControlList{})

	// Inisialisasi data
	models.SeedData(config.DB)

	// Inisialisasi Fiber
	app := fiber.New()

	// Gunakan logger untuk mencatat request
	app.Use(logger.New())

	// Middleware untuk autentikasi
	app.Use(middleware.AuthMiddleware)

	// Route untuk admin (RBAC-based)
	app.Get("/admin", middleware.RoleMiddleware("admin"), handlers.AdminHandler)

	// Route untuk editor (ACL-based)
	app.Post("/edit", middleware.ACLMiddleware, handlers.EditorHandler)

	// Route untuk viewer (ACL-based)
	app.Get("/view", middleware.ACLMiddleware, handlers.ViewerHandler)

	// Jalankan server di port yang ditentukan dalam konfigurasi
	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	logrus.Infof("Server running on port %s", port)
	log.Fatal(app.Listen(port))
}
