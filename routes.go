package main

import (
	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.Headers("Content-Type", "application/json")

	// Route for user endpooint.
	gorillaRoute.HandleFunc("/", indexEndpointHandler).Methods("GET", "POST")
	// Route for project release endpoint.
	gorillaRoute.HandleFunc("/api/releases/{user:[a-zA-Z0-9-]+}/{repo:[a-zA-Z0-9-]+}", releaseEndpointHandler)

	// User endpoints.
	// GET - Route for user endpoint.
	gorillaRoute.HandleFunc("/api/user/{user:[a-zA-Z0-9-]+}", getUserEndpointHandler).Methods("GET")
	// POST - Route for user endpoint.
	gorillaRoute.HandleFunc("/api/user/{user:[a-zA-Z0-9-]+}", postUserEndpointHandler).Methods("POST")

	return gorillaRoute
}