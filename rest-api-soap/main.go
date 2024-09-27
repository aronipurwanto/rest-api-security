package main

import (
	"log"
	"rest-api-soap/config"
)

func main() {
	// Load configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize service and pass config
	soapService := service.NewSOAPService(appConfig.SOAP)

	// Initialize Gin
	router := gin.Default()

	// Register routes and inject service to controller
	router.POST("/soap", controller.NewSOAPController(soapService).HandleSOAPRequest)

	// Start server
	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
