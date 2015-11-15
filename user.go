package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	UserName  string
	FirstName string
	Surname   string
	Role      string
	Created   int
}

func userEndpointHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	var user User
	user = getUser(userName)
	json.NewEncoder(w).Encode(user)
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
