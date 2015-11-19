package main

type User struct {
	UserName  string
	FirstName string
	Surname   string
	Role      string
	Created   int
}

func getUser(userName string) User {
	var user User

	user = User{
		UserName:  userName,
		FirstName: "Vin",
		Surname:   "Gardner",
		Role:      "God",
		Created:   12344,
	}

	return user
}

func saveUser(user User) {
	saveObject("user", user)
}
