package models

import (
	"gorm.io/gorm"
)

// User struct menyimpan informasi pengguna
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string
	OAuthID  string
}

// SeedData inisialisasi data ke database
func SeedData(db *gorm.DB) {
	// Seed data can be added here if needed
}
