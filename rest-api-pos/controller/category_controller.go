package controller

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-pos/model"
	"rest-api-pos/repository"
	"rest-api-pos/response"
)

type CategoryController struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryController(categoryRepo repository.CategoryRepository) *CategoryController {
	return &CategoryController{categoryRepo}
}

// Get All Categories
func (c *CategoryController) GetCategories(ctx *fiber.Ctx) error {
	categories, err := c.categoryRepo.GetAllCategories()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to fetch categories"))
	}
	return ctx.JSON(response.SuccessResponse(categories, "Categories fetched successfully"))
}

// Get Category by ID
func (c *CategoryController) GetCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid category ID"))
	}

	category, err := c.categoryRepo.GetCategoryByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(404, "Category not found"))
	}
	return ctx.JSON(response.SuccessResponse(category, "Category fetched successfully"))
}

// Create Category
func (c *CategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var category model.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request body"))
	}

	if err := c.categoryRepo.CreateCategory(&category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to create category"))
	}
	return ctx.Status(fiber.StatusCreated).JSON(response.SuccessResponse(category, "Category created successfully"))
}

// Update Category
func (c *CategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid category ID"))
	}

	var category model.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request body"))
	}

	// Set ID to ensure the correct category is updated
	category.ID = uint(id)

	if err := c.categoryRepo.UpdateCategory(&category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to update category"))
	}
	return ctx.JSON(response.SuccessResponse(category, "Category updated successfully"))
}

// Delete Category
func (c *CategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid category ID"))
	}

	if err := c.categoryRepo.DeleteCategory(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to delete category"))
	}
	return ctx.JSON(response.SuccessResponse(nil, "Category deleted successfully"))
}
