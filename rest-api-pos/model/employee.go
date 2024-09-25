package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string `json:"name"`
	Position string `json:"position"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	HireDate string `json:"hire_date"`
}
