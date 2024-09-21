package models

import "gorm.io/gorm"

// User struct untuk menyimpan informasi pengguna
type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"unique"`
	Role     Role   `gorm:"foreignKey:RoleID"`
	RoleID   uint
}

// Role struct menyimpan informasi role dan authorities
type Role struct {
	gorm.Model
	Name        string
	Authorities string // Authorities disimpan sebagai string (dipisah dengan koma)
}

// SeedData inisialisasi data user dan role ke database
func SeedData(db *gorm.DB) {
	// Buat roles
	adminRole := Role{Name: "admin", Authorities: "CREATE,READ,UPDATE,DELETE"}
	editorRole := Role{Name: "editor", Authorities: "READ,UPDATE"}
	viewerRole := Role{Name: "viewer", Authorities: "READ"}

	// Buat users
	adminUser := User{Name: "Admin User", Username: "admin", Role: adminRole}
	editorUser := User{Name: "Editor User", Username: "editor", Role: editorRole}
	viewerUser := User{Name: "Viewer User", Username: "viewer", Role: viewerRole}

	// Simpan ke database
	db.Create(&adminRole)
	db.Create(&editorRole)
	db.Create(&viewerRole)
	db.Create(&adminUser)
	db.Create(&editorUser)
	db.Create(&viewerUser)
}
