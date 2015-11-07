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
	"strings"
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

func releaseEndpointHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	userName := urlParams["user"]
	repositoryName := urlParams["repo"]

	var releases []Release
	releases = getReleases(userName, repositoryName)
	json.NewEncoder(w).Encode(releases)
}

func getReleases(userName string, repositoryName string) []Release {
	client := getAuthenticatedGitHubClient()
	releases := []Release{}

	tags := getRepositoryTags(client, userName, repositoryName)
	tagCounter := 0
	var lastTagName string

	for _, value := range tags {
		tagCounter++

		if tagCounter == 1 {
			lastTagName = *value.Name
			continue
		}

		//Remove this later
		if tagCounter == 6 {
			lastTagName = *value.Name
			break
		}

		releaseNames := []string{*value.Name, lastTagName}
		releaseName := strings.Join(releaseNames, "-")

		commits, err := getCommitComparison(client, userName, repositoryName, *value.Name, lastTagName)
		if err != nil {
			fmt.Println("Houston we have an error")
			continue
		}

		releases = append(releases, Release{releaseName, commits})
		lastTagName = *value.Name
	}
	return releases
}

func getAuthenticatedGitHubClient() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_CLIENT_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	return client
}

func getRepositories(client *github.Client) {
	repos, _, _ := client.Repositories.List("", nil)

	var s []string
	var owner []string
	for _, value := range repos {
		s = append(s, *value.FullName)
		owner = append(owner, *value.Owner.Login)
	}
	fmt.Println(owner)
}

func getRepositoryTags(client *github.Client, userName string, repositoryName string) []github.RepositoryTag {

	repos, _, _ := client.Repositories.ListTags(userName, repositoryName, nil)

	var s []string
	for _, value := range repos {
		s = append(s, *value.Name)
	}
	return repos
}

func getCommitComparison(client *github.Client, userName string, repositoryName string, branch1 string, branch2 string) ([]Commit, error) {
	repos, _, err := client.Repositories.CompareCommits(userName, repositoryName, branch1, branch2)

	if err != nil {
		return nil, err
	}

	var commits []Commit
	for _, value := range repos.Commits {
		message := *value.Commit.Message
		commits = append(commits, Commit{*value.SHA, message, *value.Commit.Author.Name})
	}
	return commits, nil
}

func main() {
	// Load vars.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/releases/{user:[a-zA-Z0-9-]+}/{repo:[a-zA-Z0-9-]+}", releaseEndpointHandler)
	http.Handle("/", gorillaRoute)

	serverErr := http.ListenAndServe(":3002", nil)
	if serverErr != nil {
		log.Fatal(serverErr)
	}

}
