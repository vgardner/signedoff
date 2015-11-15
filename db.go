package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Person struct {
	Name  string
	Phone string
}

func saveObject(collectionName string, userName string) {
	session := getDbSession()
	defer session.Close()

	collection := session.DB("signedoff").C(collectionName)
	err := collection.Insert(&Person{userName, "+55 53 8116 9639"})
	if err != nil {
		log.Fatal(err)
	}
}

func getObject(collectionName string, name string) interface{} {
	// Start DB session.
	session, err := mgo.Dial("localhost:4444")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C(collectionName)

	result := Person{}
	err = c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func dbTest(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	saveObject("people", userName)

	//result := getObject("people", userName)

	json.NewEncoder(w).Encode("saved")
}

func getDbSession() *mgo.Session {
	// Start session.
	session, err := mgo.Dial("localhost:4444")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func dbTestx(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]

	session, err := mgo.Dial("localhost:4444")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{userName, "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": userName}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}
