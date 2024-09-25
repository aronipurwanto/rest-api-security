package repository

import (
	"gorm.io/gorm"
	"rest-api-pos/model"
)

type SalesRepository interface {
	CreateSale(tx *gorm.DB, sale *model.Sale) error
	CreateSaleDetail(tx *gorm.DB, saleDetail *model.SaleDetail) error
	GetStockByProductID(tx *gorm.DB, productID uint) (*model.Stock, error)
	UpdateStock(tx *gorm.DB, stock *model.Stock) error
}

type salesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) SalesRepository {
	return &salesRepository{db}
}

func (r *salesRepository) CreateSale(tx *gorm.DB, sale *model.Sale) error {
	return tx.Create(&sale).Error
}

func (r *salesRepository) CreateSaleDetail(tx *gorm.DB, saleDetail *model.SaleDetail) error {
	return tx.Create(&saleDetail).Error
}

func (r *salesRepository) GetStockByProductID(tx *gorm.DB, productID uint) (*model.Stock, error) {
	var stock model.Stock
	if err := tx.Where("product_id = ?", productID).First(&stock).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *salesRepository) UpdateStock(tx *gorm.DB, stock *model.Stock) error {
	return tx.Save(&stock).Error
}
