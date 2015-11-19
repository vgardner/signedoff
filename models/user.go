package models

import "github.com/vgardner/signedoff-api/db"

type User struct {
	UserName  string
	FirstName string
	Surname   string
	Role      string
	Created   int
}

func GetUser(userName string) User {
	return User{
		UserName:  userName,
		FirstName: "Vin",
		Surname:   "Gardner",
		Role:      "God",
		Created:   12344,
	}
}

func SaveUser(user User) {
	db.SaveObject("user", user)
}
