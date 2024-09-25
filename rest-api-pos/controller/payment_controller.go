package controller

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-pos/config"
	"rest-api-pos/model"
	"rest-api-pos/repository"
	"rest-api-pos/response"
)

type PaymentController struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentController(paymentRepo repository.PaymentRepository) *PaymentController {
	return &PaymentController{paymentRepo}
}

func (p *PaymentController) CreatePayment(c *fiber.Ctx) error {
	var payment model.Payment
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(400, "Invalid request"))
	}

	// Mulai operasi transaksional
	tx := config.DB.Begin()

	// Cek apakah transaksi penjualan ada
	sale, err := p.paymentRepo.GetSaleByID(tx, payment.SaleID)
	if err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(404, "Sale not found"))
	}

	// Buat pembayaran
	if err := p.paymentRepo.CreatePayment(tx, &payment); err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to create payment"))
	}

	// Update status penjualan menjadi "completed"
	if err := p.paymentRepo.UpdateSaleStatus(tx, sale.ID, "completed"); err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(500, "Failed to update sale status"))
	}

	// Commit transaksi
	tx.Commit()

	return c.JSON(response.SuccessResponse(payment, "Payment successful"))
}
