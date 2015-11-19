package models

import "github.com/vgardner/signedoff-api/db"

// User represents a user value.
type User struct {
	UserName  string
	FirstName string
	Surname   string
	Role      string
	Created   int
}

// GetUser returns a user value.
func GetUser(userName string) User {
	return User{
		UserName:  userName,
		FirstName: "Vin",
		Surname:   "Gardner",
		Role:      "God",
		Created:   12344,
	}
}

// SaveUser persists a user value.
func SaveUser(user User) {
	db.SaveObject("user", user)
}
