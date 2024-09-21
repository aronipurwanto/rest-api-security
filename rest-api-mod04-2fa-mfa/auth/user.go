package auth

import "golang.org/x/crypto/bcrypt"

// Model untuk User
type User struct {
	Username     string
	Password     string
	TwoFAEnabled bool
	TwoFASecret  string
}

// Simulasi database pengguna
var Users = map[string]*User{}

// Fungsi untuk membuat pengguna baru
func CreateUser(username, password string) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Simpan pengguna ke "database"
	Users[username] = &User{
		Username:     username,
		Password:     string(hashedPassword),
		TwoFAEnabled: false,
	}
	return nil
}

// Fungsi untuk memverifikasi password pengguna
func VerifyPassword(username, password string) bool {
	user, exists := Users[username]
	if !exists {
		return false
	}
	// Verifikasi hash password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// Fungsi untuk mengaktifkan 2FA
func Enable2FA(username, secret string) {
	user := Users[username]
	user.TwoFAEnabled = true
	user.TwoFASecret = secret
}
