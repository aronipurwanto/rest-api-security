package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api-mod02-http-method/controllers"
	"rest-api-mod02-http-method/services"
)

func main() {
	// Initialize the service and controller
	itemService := services.ItemService{}
	itemController := controllers.ItemController{ItemService: itemService}

	// Define routes and attach the handler functions
	http.HandleFunc("/items", itemController.GetAllItemsHandler)
	http.HandleFunc("/item", itemController.GetItemHandler)
	http.HandleFunc("/item/create", itemController.CreateItemHandler)
	http.HandleFunc("/item/update", itemController.UpdateItemHandler)
	http.HandleFunc("/item/delete", itemController.DeleteItemHandler)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
