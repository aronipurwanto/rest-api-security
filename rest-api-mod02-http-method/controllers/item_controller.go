package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-mod02-http-method/models"
	"rest-api-mod02-http-method/services"
	"strconv"
)

// ItemController handles HTTP requests related to items
type ItemController struct {
	ItemService services.ItemService
}

// sendJSONResponse is a helper function to standardize the JSON response format
func sendJSONResponse(w http.ResponseWriter, code string, message string, data interface{}) {
	response := models.Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllItemsHandler handles the GET request to retrieve all items
func (c *ItemController) GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	items := c.ItemService.GetAllItems()
	sendJSONResponse(w, "200", "Items retrieved successfully", items)
}

// GetItemHandler handles the GET request to retrieve an item by ID
func (c *ItemController) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	item, found := c.ItemService.GetItemByID(id)
	if !found {
		sendJSONResponse(w, "404", "Item not found", nil)
		return
	}
	sendJSONResponse(w, "200", "Item retrieved successfully", item)
}

// CreateItemHandler handles the POST request to create a new item
func (c *ItemController) CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid request body", nil)
		return
	}
	createdItem := c.ItemService.CreateItem(newItem)
	sendJSONResponse(w, "201", "Item created successfully", createdItem)
}

// UpdateItemHandler handles the PUT request to update an item by ID
func (c *ItemController) UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	var updatedItem models.Item
	err = json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid request body", nil)
		return
	}
	item, updated := c.ItemService.UpdateItem(id, updatedItem)
	if !updated {
		sendJSONResponse(w, "404", "Item not found", nil)
		return
	}
	sendJSONResponse(w, "200", "Item updated successfully", item)
}

// DeleteItemHandler handles the DELETE request to delete an item by ID
func (c *ItemController) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	deleted := c.ItemService.DeleteItem(id)
	if !deleted {
		sendJSONResponse(w, "404", "Item not found", nil)
		return
	}
	sendJSONResponse(w, "204", "Item deleted successfully", nil)
}
