package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

type Index struct {
	Hostname    string
	UserUrl     string
	ReleasesUrl string
}

var RouteMap = map[string]http.HandlerFunc{

	"Root":      Root,
	"UserIndex": userIndex,
	"UserPost":  userPost,
	"UserGet":   userGet,
	"UserPut":   userPut,
	"DBTest":    DBTest,
}

func Root(w http.ResponseWriter, r *http.Request) {
	user := Index{
		Hostname:    baseURL(),
		UserUrl:     URL("api/user"),
		ReleasesUrl: URL("api/releases"),
	}
	json.NewEncoder(w).Encode(user)
}

func userIndex(w http.ResponseWriter, r *http.Request) {
	type userIndex struct {
		UserEndpoint string
	}
	json.NewEncoder(w).Encode(userIndex{UserEndpoint: URL("api/user/vgardner")})
}

func userPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("data"))
	userFormData := []byte(`{"UserName":"hello", "Surname":"gardner", "FirstName": "Vin", "Role": "approver"}`)
	var user User
	err := json.Unmarshal(userFormData, &user)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(user)
}

func userGet(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	var user User
	user = getUser(userName)

	json.NewEncoder(w).Encode(user)
}

func userPut(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	var user User
	user = getUser("Not saving this guy " + userName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DBTest(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	saveObject("people", &Person{userName, "+55 53 8116 9639"})

	result := getObject("people", bson.M{"name": userName})

	json.NewEncoder(w).Encode(result)
}
