package main

import "os"

func URL(path string) string {
	var URL string
	URL = baseURL() + "/" + path
	return URL
}

func baseURL() string {
	hostname, _ := os.Hostname()
	return "http://" + hostname + ":3002"
}
