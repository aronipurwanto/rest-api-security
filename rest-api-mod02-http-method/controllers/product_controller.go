package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-mod02-http-method/models"
	"rest-api-mod02-http-method/services"
	"strconv"
)

// ProductController handles HTTP requests related to Products
type ProductController struct {
	ProductService services.ProductService
}

// GetAllProductsHandler handles the GET request to retrieve all Products
func (c *ProductController) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	Products := c.ProductService.GetAllProducts()
	sendJSONResponse(w, "200", "Products retrieved successfully", Products)
}

// GetProductHandler handles the GET request to retrieve an Product by ID
func (c *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	Product, found := c.ProductService.GetProductByID(id)
	if !found {
		sendJSONResponse(w, "404", "Product not found", nil)
		return
	}
	sendJSONResponse(w, "200", "Product retrieved successfully", Product)
}

// CreateProductHandler handles the POST request to create a new Product
func (c *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid request body", nil)
		return
	}
	createdProduct := c.ProductService.CreateProduct(newProduct)
	sendJSONResponse(w, "201", "Product created successfully", createdProduct)
}

// UpdateProductHandler handles the PUT request to update an Product by ID
func (c *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid request body", nil)
		return
	}
	Product, updated := c.ProductService.UpdateProduct(id, updatedProduct)
	if !updated {
		sendJSONResponse(w, "404", "Product not found", nil)
		return
	}
	sendJSONResponse(w, "200", "Product updated successfully", Product)
}

// DeleteProductHandler handles the DELETE request to delete an Product by ID
func (c *ProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendJSONResponse(w, "400", "Invalid ID", nil)
		return
	}
	deleted := c.ProductService.DeleteProduct(id)
	if !deleted {
		sendJSONResponse(w, "404", "Product not found", nil)
		return
	}
	sendJSONResponse(w, "204", "Product deleted successfully", nil)
}
