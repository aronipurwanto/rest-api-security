package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"rest-api-mod05-jwt/config"
	"rest-api-mod05-jwt/models"
	"time"
)

// Secret key for signing JWT
var jwtSecret = []byte(viper.GetString("jwt.secret"))

// Claims struct untuk JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// HashPassword melakukan hash pada password pengguna
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash untuk memverifikasi password pengguna
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Authenticate untuk melakukan autentikasi pengguna berdasarkan username dan password
func Authenticate(username, password string) (models.User, error) {
	var user models.User
	result := config.DB.Preload("Role").Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	if !CheckPasswordHash(password, user.Password) {
		return user, result.Error
	}

	return user, nil
}

// GenerateJWT untuk membuat JWT setelah login berhasil
func GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(time.Duration(viper.GetInt("jwt.expires_in_minutes")) * time.Minute)

	claims := &Claims{
		Username: user.Username,
		Role:     user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
