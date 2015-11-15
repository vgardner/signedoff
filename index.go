package main

import (
	"encoding/json"
	"net/http"
)

type Index struct {
	User     string
	Releases string
	Surname  string
	Role     string
	Created  int
}

func indexEndpointHandler(w http.ResponseWriter, r *http.Request) {
	var user Index
	user = getIndex()
	w.Header().Set("Content-Type", "application/json")
	//.MarshalIndent(data, "", "\t")
	json.NewEncoder(w).Encode(user)
}

func getIndex() Index {
	var user Index

	user = Index{
		User:     "Vin",
		Releases: "Gardner",
		Role:     "God",
		Created:  12344,
	}

	return user
}
