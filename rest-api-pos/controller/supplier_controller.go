package controller

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-pos/model"
	"rest-api-pos/repository"
	"rest-api-pos/response"
)

type SupplierController struct {
	supplierRepo repository.SupplierRepository
}

func NewSupplierController(supplierRepo repository.SupplierRepository) *SupplierController {
	return &SupplierController{supplierRepo}
}

// Get All Suppliers
func (s *SupplierController) GetSuppliers(c *fiber.Ctx) error {
	suppliers, err := s.supplierRepo.GetAllSuppliers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to fetch suppliers"))
	}
	return c.JSON(response.SuccessResponse(suppliers, "Suppliers fetched successfully"))
}

// Get Supplier by ID
func (s *SupplierController) GetSupplier(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid supplier ID"))
	}

	supplier, err := s.supplierRepo.GetSupplierByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(404, "Supplier not found"))
	}
	return c.JSON(response.SuccessResponse(supplier, "Supplier fetched successfully"))
}

// Create Supplier
func (s *SupplierController) CreateSupplier(c *fiber.Ctx) error {
	var supplier model.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request body"))
	}

	if err := s.supplierRepo.CreateSupplier(&supplier); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to create supplier"))
	}
	return c.Status(fiber.StatusCreated).JSON(response.SuccessResponse(supplier, "Supplier created successfully"))
}

// Update Supplier
func (s *SupplierController) UpdateSupplier(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid supplier ID"))
	}

	var supplier model.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request body"))
	}

	// Set ID to ensure the correct supplier is updated
	supplier.ID = uint(id)

	if err := s.supplierRepo.UpdateSupplier(&supplier); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to update supplier"))
	}
	return c.JSON(response.SuccessResponse(supplier, "Supplier updated successfully"))
}

// Delete Supplier
func (s *SupplierController) DeleteSupplier(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid supplier ID"))
	}

	if err := s.supplierRepo.DeleteSupplier(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to delete supplier"))
	}
	return c.JSON(response.SuccessResponse(nil, "Supplier deleted successfully"))
}
