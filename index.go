package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Index struct {
	Hostname    string
	UserUrl     string
	ReleasesUrl string
}

var RouteMap = map[string]http.HandlerFunc{

	"Root": IndexEndpointHandler,
}

func IndexEndpointHandler(w http.ResponseWriter, r *http.Request) {
	var user Index
	user = getIndex()
	w.Header().Set("Content-Type", "application/json")
	//.MarshalIndent(data, "", "\t")
	json.NewEncoder(w).Encode(user)
}

func getIndex() Index {
	var user Index

	user = Index{
		Hostname:    baseUrl(),
		UserUrl:     url("api/user"),
		ReleasesUrl: url("api/releases"),
	}

	return user
}

func url(path string) string {
	var url string
	url = baseUrl() + "/" + path
	return url
}

func baseUrl() string {
	hostname, _ := os.Hostname()
	return "http://" + hostname + ":3002"
}
