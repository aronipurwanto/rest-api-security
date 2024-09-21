package models

import "gorm.io/gorm"

// User struct menyimpan informasi pengguna
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	OAuthID  string
	RoleID   uint
	Role     Role
}

// Role struct menyimpan informasi peran pengguna
type Role struct {
	gorm.Model
	Name string
}

// SeedData inisialisasi role ke database
func SeedData(db *gorm.DB) {
	adminRole := Role{Name: "admin"}
	userRole := Role{Name: "user"}

	db.Create(&adminRole)
	db.Create(&userRole)
}
