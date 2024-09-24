package models

// User model
type User struct {
	Username string
	Password string
}

// In-memory list of users
var users []User

// InitializeUsers creates sample users for login
func InitializeUsers() {
	users = []User{
		{Username: "admin", Password: "password123"},
		{Username: "user1", Password: "user1pass"},
		{Username: "user2", Password: "user2pass"},
		{Username: "manager", Password: "managerpass"},
		{Username: "guest", Password: "guestpass"},
	}
}

// ValidateUser checks if the username and password are valid
func ValidateUser(username, password string) bool {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
