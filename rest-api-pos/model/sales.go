package model

import (
	"gorm.io/gorm"
	"time"
)

type Sale struct {
	gorm.Model
	TransactionDate time.Time    `json:"transaction_date"`
	TotalAmount     float64      `json:"total_amount"`
	CustomerID      uint         `json:"customer_id"`
	Status          string       `json:"status"` // e.g., "completed", "pending"
	Customer        Customer     `gorm:"foreignKey:CustomerID"`
	SaleDetails     []SaleDetail `gorm:"foreignKey:SaleID"`
}

type SaleDetail struct {
	gorm.Model
	SaleID    uint    `json:"sale_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
