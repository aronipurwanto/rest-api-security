package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

// EncryptAES encrypts a given string using AES-GCM
func EncryptAES(key []byte, text string) ([]byte, error) {
	log.Println("Starting AES encryption process for text:", text)

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error creating AES cipher block:", err)
		return nil, err
	}

	// Generate GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Error generating GCM mode:", err)
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	// Encrypt the text using AES-GCM
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)

	log.Println("AES encryption successful. Ciphertext generated.")
	return ciphertext, nil
}

// DecryptAES decrypts an AES-GCM encrypted string
func DecryptAES(key []byte, ciphertext []byte) (string, error) {
	log.Println("Starting AES decryption process.")

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error creating AES cipher block:", err)
		return "", err
	}

	// Generate GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Error generating GCM mode:", err)
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the text
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println("Error during AES decryption:", err)
		return "", err
	}

	log.Println("AES decryption successful. Plaintext recovered.")
	return string(plaintext), nil
}
