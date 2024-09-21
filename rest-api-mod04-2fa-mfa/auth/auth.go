package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
)

// Fungsi untuk menghandle registrasi pengguna baru
func RegisterHandler(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Membuat pengguna baru
	if err := CreateUser(req.Username, req.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

// Fungsi untuk menghandle login
func LoginHandler(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Verifikasi password
	if !VerifyPassword(req.Username, req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	user := Users[req.Username]

	// Jika 2FA belum diaktifkan, aktifkan sekarang
	if !user.TwoFAEnabled {
		secret := GenerateTOTPSecret(req.Username)
		Enable2FA(req.Username, secret)

		// Kirim URL otentikasi OTP (QR Code URL) untuk aplikasi Google Authenticator
		otpURL := GetTOTPProvisioningURL(req.Username, secret)
		return c.JSON(fiber.Map{
			"message":      "2FA enabled. Please scan the QR code with your authenticator app.",
			"otp_auth_url": otpURL,
		})
	}

	// Jika 2FA diaktifkan, minta kode OTP
	return c.JSON(fiber.Map{"message": "2FA required"})
}

// Middleware untuk verifikasi kode OTP (2FA)
func TwoFAMiddleware(c *fiber.Ctx) error {
	username := c.Query("username")
	otpCode := c.Query("otp")

	if username == "" || otpCode == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "2FA required"})
	}

	user, exists := Users[username]
	if !exists || !user.TwoFAEnabled {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user or 2FA not enabled"})
	}

	// Verifikasi kode OTP
	valid := totp.Validate(otpCode, user.TwoFASecret)
	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OTP code"})
	}

	// Lanjutkan ke handler berikutnya jika OTP valid
	return c.Next()
}

// Fungsi untuk menangani request ke data yang diamankan
func SecureDataHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Access granted to secure data!"})
}

// Fungsi untuk menghasilkan secret TOTP
func GenerateTOTPSecret(username string) string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyAPI",
		AccountName: username,
	})
	if err != nil {
		panic(err)
	}
	return key.Secret()
}

// Fungsi untuk mendapatkan URL provisioning TOTP (untuk QR Code)
func GetTOTPProvisioningURL(username, secret string) string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyAPI",
		AccountName: username,
		Period:      30,
		SecretSize:  32,
	})
	if err != nil {
		panic(err)
	}
	return key.URL()
}
