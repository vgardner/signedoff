package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DBNAME = "signedoff"
const DBHOST = "mongo:27017"

type Person struct {
	Name  string
	Phone string
}

func SaveObject(collectionName string, object interface{}) {
	session := getDBSession()
	defer session.Close()

	collection := session.DB(DBNAME).C(collectionName)
	err := collection.Insert(object)
	if err != nil {
		log.Fatal(err)
	}
}

func GetObject(collectionName string, bsonObject bson.M) interface{} {
	session := getDBSession()
	defer session.Close()

	collection := session.DB(DBNAME).C(collectionName)

	var result interface{}
	err := collection.Find(bsonObject).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getDBSession() *mgo.Session {
	session, err := mgo.Dial(DBHOST)
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}
