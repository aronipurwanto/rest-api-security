package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"strconv"
	"time"
	"your_project/models"
)

// Kunci rahasia untuk menandatangani JWT token
var jwtSecret = []byte("your_secret_key")

// Login function untuk memulai proses login dan mengirim OTP
func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Cari user berdasarkan username dan password
	var user *models.User
	for i, u := range models.Users {
		if u.Username == username && u.Password == password {
			user = &models.Users[i]
			break
		}
	}

	// Jika user tidak ditemukan, log dan kembalikan error
	if user == nil {
		log.Println("Login gagal: Pengguna tidak ditemukan atau password salah")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Generate OTP dan simpan di user
	otp := generateOTP()
	user.OTP = otp
	user.OTPExpiry = time.Now().Add(5 * time.Minute) // OTP berlaku selama 5 menit

	log.Printf("Login berhasil: OTP telah dikirim untuk pengguna %s", user.Username)

	// Tampilkan log OTP (dalam aplikasi nyata OTP ini dikirimkan via email atau SMS)
	log.Printf("OTP untuk %s: %s", user.Username, user.OTP)

	return c.JSON(fiber.Map{
		"message": "OTP sent to your email",
	})
}

// VerifyOTP function untuk memverifikasi OTP yang dimasukkan pengguna
func VerifyOTP(c *fiber.Ctx) error {
	username := c.FormValue("username")
	otp := c.FormValue("otp")

	// Cari user berdasarkan username
	var user *models.User
	for i, u := range models.Users {
		if u.Username == username {
			user = &models.Users[i]
			break
		}
	}

	// Jika user tidak ditemukan atau OTP sudah expired, log dan kembalikan error
	if user == nil || time.Now().After(user.OTPExpiry) {
		log.Println("Verifikasi OTP gagal: Pengguna tidak ditemukan atau OTP kadaluarsa")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or expired OTP",
		})
	}

	// Verifikasi OTP
	if user.OTP != otp {
		log.Println("Verifikasi OTP gagal: OTP tidak valid")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid OTP",
		})
	}

	log.Printf("Verifikasi OTP berhasil untuk pengguna %s", user.Username)

	// Generate JWT token setelah verifikasi OTP berhasil
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
	})

	// Tanda tangani token dengan kunci rahasia
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Error saat menghasilkan JWT token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}

	log.Printf("JWT token dihasilkan untuk pengguna %s", user.Username)

	// Kembalikan token ke klien
	return c.JSON(fiber.Map{
		"message": "Authentication successful",
		"token":   tokenString,
	})
}

// generateOTP function untuk membuat OTP acak
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(1000 + rand.Intn(9000)) // Menghasilkan OTP 4 digit antara 1000-9999
}
