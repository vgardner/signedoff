package main

import (
	"encoding/json"
	"net/http"
)

type Index struct {
	UserName  string
	FirstName string
	Surname   string
	Role      string
	Created   int
}

func indexEndpointHandler(w http.ResponseWriter, r *http.Request) {
	var user Index
	user = getIndex()
	json.NewEncoder(w).Encode(user)
}

func getIndex() Index {
	var user Index

	user = Index{
		FirstName: "Vin",
		Surname:   "Gardner",
		Role:      "God",
		Created:   12344,
	}

	return user
}
