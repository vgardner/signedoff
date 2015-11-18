package main

import (
	"log"
	"net/http"

	"github.com/EconomistDigitalSolutions/ramlapi"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var router *mux.Router

func init() {
	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	NewRouter()
	log.Println("running...")
	log.Fatal(http.ListenAndServe(":3002", nil))
}

func NewRouter() *mux.Router {
	router = mux.NewRouter().StrictSlash(true)
	assembleMiddleware()
	assembleRoutes(router)
	return router
}

// assembleMiddleware sets up the middleware stack.
func assembleMiddleware() {
	http.Handle("/", JSONMiddleware(router))
}

func assembleRoutes(r *mux.Router) {
	api, err := ramlapi.Process("api.raml")
	if err != nil {
		log.Fatal(err)
	}
	ramlapi.Build(api, routerFunc)
}

func routerFunc(ep *ramlapi.Endpoint) {
	path := ep.Path

	router.
		Methods(ep.Verb).
		Path(path).
		Handler(RouteMap[ep.Handler])
}
