package services

import (
	"rest-api-mod05-auth/config"
	"rest-api-mod05-auth/models"
	"strings"
)

// Authenticate user berdasarkan username dari database
func Authenticate(username string) (models.User, bool) {
	var user models.User
	result := config.DB.Preload("Role").Where("username = ?", username).First(&user)
	return user, result.Error == nil
}

// Check apakah user memiliki authority tertentu
func HasAuthority(user models.User, authority string) bool {
	authorities := strings.Split(user.Role.Authorities, ",")
	for _, auth := range authorities {
		if strings.EqualFold(auth, authority) {
			return true
		}
	}
	return false
}
