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

func getUserEndpointHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	var user User
	user = getUser(userName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func postUserEndpointHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	var user User
	user = getUser("Not saving this guy " + userName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// User index endpoint handler.
func userIndexEndpointHandler(w http.ResponseWriter, r *http.Request) {
	type userIndex struct {
		UserEndpoint string
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userIndex{UserEndpoint: url("api/user/vgardner")})
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
