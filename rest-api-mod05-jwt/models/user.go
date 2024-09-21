package models

import "gorm.io/gorm"

// User struct menyimpan informasi pengguna
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	RoleID   uint
	Role     Role
}

// Role struct menyimpan informasi peran pengguna
type Role struct {
	gorm.Model
	Name string
}

// SeedData untuk inisialisasi role dan user
func SeedData(db *gorm.DB) {
	// Buat roles
	adminRole := Role{Name: "admin"}
	userRole := Role{Name: "user"}

	// Simpan roles ke database
	db.Create(&adminRole)
	db.Create(&userRole)

	// Buat users
	adminUser := User{Username: "admin", Password: "adminpassword", Role: adminRole}
	normalUser := User{Username: "user", Password: "userpassword", Role: userRole}

	// Simpan users ke database
	db.Create(&adminUser)
	db.Create(&normalUser)
}
