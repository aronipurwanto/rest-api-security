package controller

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-pos/model"
	"rest-api-pos/repository"
	"rest-api-pos/response"
)

type ProductController struct {
	productRepo repository.ProductRepository
}

func NewProductController(productRepo repository.ProductRepository) *ProductController {
	return &ProductController{productRepo}
}

// Get All Products
func (p *ProductController) GetProducts(c *fiber.Ctx) error {
	products, err := p.productRepo.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to fetch products"))
	}
	return c.JSON(response.SuccessResponse(products, "Products fetched successfully"))
}

// Get Product by ID
func (p *ProductController) GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid product ID"))
	}

	product, err := p.productRepo.GetProductByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(404, "Product not found"))
	}
	return c.JSON(response.SuccessResponse(product, "Product fetched successfully"))
}

// Create Product
func (p *ProductController) CreateProduct(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request body"))
	}

	if err := p.productRepo.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to create product"))
	}
	return c.Status(fiber.StatusCreated).JSON(response.SuccessResponse(product, "Product created successfully"))
}

// Update Product
func (p *ProductController) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid product ID"))
	}

	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request body"))
	}

	// Set ID to ensure the correct product is updated
	product.ID = uint(id)

	if err := p.productRepo.UpdateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to update product"))
	}
	return c.JSON(response.SuccessResponse(product, "Product updated successfully"))
}

// Delete Product
func (p *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid product ID"))
	}

	if err := p.productRepo.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to delete product"))
	}
	return c.JSON(response.SuccessResponse(nil, "Product deleted successfully"))
}
