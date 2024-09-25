package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	SaleID        uint    `json:"sale_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	PaymentDate   string  `json:"payment_date"`
	Sale          Sale    `gorm:"foreignKey:SaleID"`
}
