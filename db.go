package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

const DBNAME = "signedoff"
const DBHOST = "localhost:4444"

type Person struct {
	Name  string
	Phone string
}

func saveObject(collectionName string, object interface{}) {
	session := getDbSession()
	defer session.Close()

	collection := session.DB(DBNAME).C(collectionName)
	err := collection.Insert(object)
	if err != nil {
		log.Fatal(err)
	}
}

func getObject(collectionName string, name string) interface{} {
	session := getDbSession()
	defer session.Close()

	collection := session.DB(DBNAME).C(collectionName)

	var result interface{}
	err := collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func dbTest(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	saveObject("people", &Person{userName, "+55 53 8116 9639"})

	result := getObject("people", userName)

	json.NewEncoder(w).Encode(result)
}

func getDbSession() *mgo.Session {
	// Start session.
	session, err := mgo.Dial(DBHOST)
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}
