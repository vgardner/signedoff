package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	name := urlParams["user"]
	HelloMessage := "Hello, " + name

	message := API{HelloMessage}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output))
}

func getRepositoryInfo(w http.ResponseWriter, r *http.Request) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_CLIENT_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	//repos, _, err := client.Repositories.List("", nil)
	orgs, _, _ := client.Organizations.List("vgardner", nil)
	fmt.Fprintf(w, *orgs[1].Login)
}

func main() {
	// Load vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	gorillaRoute.HandleFunc("/api/repo/{repo:[a-z]+}", getRepositoryInfo)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3001", nil)
}
