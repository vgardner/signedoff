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
	fmt.Fprintf(w, "test")
}

func getAuthenticatedGitHubClient() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_CLIENT_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	return client
}

func getRepositoryData(client *github.Client) {

	// list all repositories for the authenticated user
	repos, _, _ := client.Repositories.List("", nil)
	//orgs, _, _ := client.Organizations.List("vgardner", nil)

	var s []string
	var owner []string
	for _, value := range repos {
		s = append(s, *value.FullName)
		owner = append(owner, *value.Owner.Login)
	}
	fmt.Println(s)
	fmt.Println(owner)
}

func getRepositoryTags(client *github.Client) {
	//list all repositories for the authenticated user
	repos, _, _ := client.Repositories.ListTags("EconomistDigitalSolutions", "website", nil)

	var s []string
	for _, value := range repos {
		s = append(s, *value.Name)
	}
	fmt.Println(s)
}

func getCommitComparison(client *github.Client) []github.RepositoryCommit {
	//list all repositories for the authenticated user
	repos, _, _ := client.Repositories.CompareCommits("EconomistDigitalSolutions", "website", "release-296.0", "release-297.0")

	var s []string
	for _, value := range repos.Commits {
		s = append(s, *value.SHA)
	}
	return repos.Commits
}

type Release struct {
	releaseId string
	commits   []github.RepositoryCommit
}

func main() {
	// Load vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var release = Release{}

	client := getAuthenticatedGitHubClient()
	getRepositoryData(client)

	getRepositoryTags(client)

	release.commits = getCommitComparison(client)

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	gorillaRoute.HandleFunc("/api/repo/{repo:[a-z]+}", getRepositoryInfo)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3001", nil)
}
