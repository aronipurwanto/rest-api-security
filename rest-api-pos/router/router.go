package router

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-pos/config"
	"rest-api-pos/controller"
	"rest-api-pos/middleware"
	"rest-api-pos/model"
	"rest-api-pos/repository"
)

func SetupRoutes(app *fiber.App) {
	// Repository initialization
	salesRepo := repository.NewSalesRepository(config.DB)
	paymentRepo := repository.NewPaymentRepository(config.DB)
	supplierRepo := repository.NewSupplierRepository(config.DB)
	productRepo := repository.NewProductRepository(config.DB)
	categoryRepo := repository.NewCategoryRepository(config.DB)

	// Controller initialization
	salesController := controller.NewSalesController(salesRepo)
	paymentController := controller.NewPaymentController(paymentRepo)
	supplierController := controller.NewSupplierController(supplierRepo)
	productController := controller.NewProductController(productRepo)
	categoryController := controller.NewCategoryController(categoryRepo)

	// Auth Routes
	app.Post("/login", controller.Login)
	app.Post("/refresh-token", controller.RefreshToken)
	app.Get("/profile", middleware.AuthMiddleware, controller.GetProfile)

	// Supplier Routes
	app.Get("/suppliers", middleware.AuthMiddleware, supplierController.GetSuppliers)
	app.Get("/suppliers/:id", middleware.AuthMiddleware, supplierController.GetSupplier)
	app.Post("/suppliers", middleware.AuthMiddleware, supplierController.CreateSupplier)
	app.Put("/suppliers/:id", middleware.AuthMiddleware, supplierController.UpdateSupplier)
	app.Delete("/suppliers/:id", middleware.AuthMiddleware, supplierController.DeleteSupplier)

	// Product Routes
	app.Get("/products", middleware.AuthMiddleware, productController.GetProducts)
	app.Get("/products/:id", middleware.AuthMiddleware, productController.GetProduct)
	app.Post("/products", middleware.AuthMiddleware, productController.CreateProduct)
	app.Put("/products/:id", middleware.AuthMiddleware, productController.UpdateProduct)
	app.Delete("/products/:id", middleware.AuthMiddleware, productController.DeleteProduct)

	// Category Routes
	app.Get("/categories", middleware.AuthMiddleware, categoryController.GetCategories)
	app.Get("/categories/:id", middleware.AuthMiddleware, categoryController.GetCategory)
	app.Post("/categories", middleware.AuthMiddleware, categoryController.CreateCategory)
	app.Put("/categories/:id", middleware.AuthMiddleware, categoryController.UpdateCategory)
	app.Delete("/categories/:id", middleware.AuthMiddleware, categoryController.DeleteCategory)

	// Sales Routes
	app.Post("/sales", middleware.AuthMiddleware, salesController.CreateSale)

	// Payment Routes
	app.Post("/payments", middleware.AuthMiddleware, paymentController.CreatePayment)

	// generate data
	model.SeedData()
}
