package model

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	LastUpdate string  `json:"last_update"`
	Product    Product `gorm:"foreignKey:ProductID"`
}
