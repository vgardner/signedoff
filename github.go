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
		message := *value.Commit.Message
		commits = append(commits, Commit{*value.SHA, message[0:20], *value.Commit.Author.Name})
	}
	return commits
}
