package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

type Release struct {
	ReleaseId string
	Commits   []Commit
}

type Commit struct {
	Sha     string
	Message string
	Author  string
}

func init() {
	// Load vars.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/releases/{user:[a-zA-Z0-9-]+}/{repo:[a-zA-Z0-9-]+}", releaseEndpointHandler)
	http.Handle("/", gorillaRoute)

	serverErr := http.ListenAndServe(":3002", nil)
	if serverErr != nil {
		log.Fatal(serverErr)
	}

}

func releaseEndpointHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]
	repositoryName := urlParams["repo"]

	var releases []Release
	releases = getReleases(userName, repositoryName)
	json.NewEncoder(w).Encode(releases)
}
