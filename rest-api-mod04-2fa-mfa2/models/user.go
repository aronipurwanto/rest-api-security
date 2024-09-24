package models

import "time"

// User struct untuk menyimpan data pengguna
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	OTP       string    `json:"otp,omitempty"`
	OTPExpiry time.Time `json:"otp_expiry,omitempty"`
}

// Users slice untuk menyimpan pengguna secara in-memory
var Users = []User{
	{ID: 1, Username: "user1", Password: "password123"},
	{ID: 2, Username: "user2", Password: "password456"},
}
