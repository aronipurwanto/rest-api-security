package repository

import (
	"gorm.io/gorm"
	"rest-api-pos/model"
)

type PaymentRepository interface {
	CreatePayment(tx *gorm.DB, payment *model.Payment) error
	UpdateSaleStatus(tx *gorm.DB, saleID uint, status string) error
	GetSaleByID(tx *gorm.DB, saleID uint) (*model.Sale, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) CreatePayment(tx *gorm.DB, payment *model.Payment) error {
	return tx.Create(&payment).Error
}

func (r *paymentRepository) UpdateSaleStatus(tx *gorm.DB, saleID uint, status string) error {
	return tx.Model(&model.Sale{}).Where("id = ?", saleID).Update("status", status).Error
}

func (r *paymentRepository) GetSaleByID(tx *gorm.DB, saleID uint) (*model.Sale, error) {
	var sale model.Sale
	if err := tx.First(&sale, saleID).Error; err != nil {
		return nil, err
	}
	return &sale, nil
}
