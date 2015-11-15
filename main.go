package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Start router.
	http.Handle("/", getRouter())

	serverErr := http.ListenAndServe(":3002", nil)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
