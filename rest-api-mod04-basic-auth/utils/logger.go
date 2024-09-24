package utils

import "log"

// Custom logging utility (optional)
func LogInfo(message string) {
	log.Println("[INFO] " + message)
}

func LogError(message string) {
	log.Println("[ERROR] " + message)
}
