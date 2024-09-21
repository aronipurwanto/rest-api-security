package services

import (
	"rest-api-mod05-rbac/config"
	"rest-api-mod05-rbac/models"
)

// Authenticate user berdasarkan username dari database
func Authenticate(username string) (models.User, bool) {
	var user models.User
	result := config.DB.Preload("Role").Where("username = ?", username).First(&user)
	return user, result.Error == nil
}

// Check apakah user memiliki izin untuk akses ACL
func HasAccessControl(user models.User, path, method string) bool {
	var acl models.AccessControlList
	result := config.DB.Where("user_id = ? AND path = ? AND method = ?", user.ID, path, method).First(&acl)
	return result.Error == nil
}
