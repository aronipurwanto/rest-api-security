package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	SupplierID   uint          `json:"supplier_id"`
	OrderDate    string        `json:"order_date"`
	TotalAmount  float64       `json:"total_amount"`
	Supplier     Supplier      `gorm:"foreignKey:SupplierID"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
