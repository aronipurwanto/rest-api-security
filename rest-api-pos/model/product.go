package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	Stock      int      `json:"stock"`
	CategoryID uint     `json:"category_id"`
	SupplierID uint     `json:"supplier_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID"`
}

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

type Supplier struct {
	gorm.Model
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Products []Product `gorm:"foreignKey:SupplierID"`
}
