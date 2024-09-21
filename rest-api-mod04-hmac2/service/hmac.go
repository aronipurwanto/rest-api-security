package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Secret key for generating HMAC
const SecretKey string = "supersecretkey"

// GenerateHMAC creates an HMAC for a given message using a secret key
func GenerateHMAC(message string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC checks if the provided HMAC matches the generated HMAC for a message
func VerifyHMAC(message, receivedMAC, key string) bool {
	expectedMAC := GenerateHMAC(message, key)
	return hmac.Equal([]byte(receivedMAC), []byte(expectedMAC))
}
