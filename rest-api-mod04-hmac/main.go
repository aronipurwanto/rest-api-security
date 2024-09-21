package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Fungsi untuk menghasilkan HMAC dari sebuah pesan
func generateHMAC(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// Fungsi untuk memverifikasi HMAC
func verifyHMAC(message, receivedMAC, secret string) bool {
	generatedMAC := generateHMAC(message, secret)
	return hmac.Equal([]byte(receivedMAC), []byte(generatedMAC))
}

func main() {
	message := "Hello, HMAC!"
	secret := "supersecretkey"

	// Menghasilkan HMAC
	mac := generateHMAC(message, secret)
	fmt.Println("Generated HMAC:", mac)

	// Verifikasi HMAC
	isValid := verifyHMAC(message, mac, secret)
	fmt.Println("Is HMAC valid?", isValid)

	// Coba dengan pesan yang berbeda
	isValid = verifyHMAC("Different message", mac, secret)
	fmt.Println("Is HMAC valid with different message?", isValid)
}
