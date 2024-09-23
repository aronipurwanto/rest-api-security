package services

import "rest-api-mod02-http-method/models"

// In-memory database
var Products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 15000, Description: "Product Laptop", Category: "Electronik"},
	{ID: 2, Name: "Smartphone", Price: 7000, Description: "Product Smartphone", Category: "Electronik"},
}

// ProductService handles the business logic for Products
type ProductService struct{}

// GetAllProducts returns all Products
func (s *ProductService) GetAllProducts() []models.Product {
	return Products
}

// GetProductByID returns an Product by its ID
func (s *ProductService) GetProductByID(id int) (models.Product, bool) {
	for _, Product := range Products {
		if Product.ID == id {
			return Product, true
		}
	}
	return models.Product{}, false
}

// CreateProduct adds a new Product
func (s *ProductService) CreateProduct(newProduct models.Product) models.Product {
	newProduct.ID = len(Products) + 1
	Products = append(Products, newProduct)
	return newProduct
}

// UpdateProduct updates an existing Product
func (s *ProductService) UpdateProduct(id int, updatedProduct models.Product) (models.Product, bool) {
	for i, Product := range Products {
		if Product.ID == id {
			updatedProduct.ID = id
			Products[i] = updatedProduct
			return updatedProduct, true
		}
	}
	return models.Product{}, false
}

// DeleteProduct deletes an Product by ID
func (s *ProductService) DeleteProduct(id int) bool {
	for i, Product := range Products {
		if Product.ID == id {
			Products = append(Products[:i], Products[i+1:]...)
			return true
		}
	}
	return false
}
