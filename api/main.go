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

type Release struct {
	ReleaseId string
	Commits   []Commit
}

type Commit struct {
	Sha      string
	Message  string
	Commiter string
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

	client := getAuthenticatedGitHubClient()

	var release = Release{}
	release.ReleaseId = "release1"
	release.Commits = getCommitComparison(client)
	//fmt.Println(release.Commits)
	json.NewEncoder(w).Encode(release.Commits)
	//fmt.Fprintf(w, release.commits[4].sha)
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

func getCommitComparison(client *github.Client) []Commit {
	//list all repositories for the authenticated user
	repos, _, _ := client.Repositories.CompareCommits("EconomistDigitalSolutions", "website", "release-296.0", "release-297.0")

	var commits []Commit
	for _, value := range repos.Commits {
		commits = append(commits, Commit{*value.SHA, "test", "hello"})
	}
	commits = append(commits, Commit{"testing", "test", "hello"})
	return commits
}

func main() {
	// Load vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := getAuthenticatedGitHubClient()
	getRepositoryData(client)

	getRepositoryTags(client)

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	gorillaRoute.HandleFunc("/api/repo/{repo:[a-z]+}", getRepositoryInfo)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3001", nil)
}
