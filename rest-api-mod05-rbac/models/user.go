package models

import "gorm.io/gorm"

// User struct menyimpan informasi pengguna
type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"unique"`
	RoleID   uint
	Role     Role
	ACL      []AccessControlList `gorm:"foreignKey:UserID"`
}

// Role struct menyimpan informasi peran pengguna (RBAC)
type Role struct {
	gorm.Model
	Name string
}

// AccessControlList (ACL) untuk menyimpan izin akses pengguna
type AccessControlList struct {
	gorm.Model
	UserID uint
	Path   string // API path
	Method string // HTTP method (GET, POST, etc.)
}

// SeedData inisialisasi data user, role, dan ACL ke database
func SeedData(db *gorm.DB) {
	// Buat roles
	adminRole := Role{Name: "admin"}
	editorRole := Role{Name: "editor"}
	viewerRole := Role{Name: "viewer"}

	// Simpan roles ke database
	db.Create(&adminRole)
	db.Create(&editorRole)
	db.Create(&viewerRole)

	// Buat users
	adminUser := User{Name: "Admin User", Username: "admin", Role: adminRole}
	editorUser := User{Name: "Editor User", Username: "editor", Role: editorRole}
	viewerUser := User{Name: "Viewer User", Username: "viewer", Role: viewerRole}

	// Simpan users ke database
	db.Create(&adminUser)
	db.Create(&editorUser)
	db.Create(&viewerUser)

	// Buat ACL untuk admin
	aclAdmin := AccessControlList{UserID: adminUser.ID, Path: "/admin", Method: "GET"}
	db.Create(&aclAdmin)

	// Buat ACL untuk editor
	aclEditor := AccessControlList{UserID: editorUser.ID, Path: "/edit", Method: "POST"}
	db.Create(&aclEditor)

	// Buat ACL untuk viewer
	aclViewer := AccessControlList{UserID: viewerUser.ID, Path: "/view", Method: "GET"}
	db.Create(&aclViewer)
}
