package controller

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-pos/config"
	"rest-api-pos/model"
	"rest-api-pos/repository"
	"rest-api-pos/response"
)

type SalesController struct {
	salesRepo repository.SalesRepository
}

func NewSalesController(salesRepo repository.SalesRepository) *SalesController {
	return &SalesController{salesRepo}
}

func (s *SalesController) CreateSale(c *fiber.Ctx) error {
	var sale model.Sale
	if err := c.BodyParser(&sale); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request"))
	}

	// Mulai operasi transaksional
	tx := config.DB.Begin()

	// Buat penjualan
	if err := s.salesRepo.CreateSale(tx, &sale); err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to create sale"))
	}

	// Buat detail penjualan dan kurangi stok
	for _, detail := range sale.SaleDetails {
		detail.SaleID = sale.ID
		if err := s.salesRepo.CreateSaleDetail(tx, &detail); err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to create sale details"))
		}

		// Kurangi stok produk
		stock, err := s.salesRepo.GetStockByProductID(tx, detail.ProductID)
		if err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Product stock not found"))
		}

		if stock.Quantity < detail.Quantity {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Insufficient stock"))
		}

		stock.Quantity -= detail.Quantity
		if err := s.salesRepo.UpdateStock(tx, stock); err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to update stock"))
		}
	}

	// Commit transaksi
	tx.Commit()

	return c.JSON(response.SuccessResponse(sale, "Sale created successfully"))
}
