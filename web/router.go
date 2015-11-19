package web

import (
	"log"
	"net/http"

	"github.com/EconomistDigitalSolutions/ramlapi"
	"github.com/gorilla/mux"
)

var router *mux.Router

func NewRouter() {
	router = mux.NewRouter().StrictSlash(true)
	assembleMiddleware()
	assembleRoutes()
}

// assembleMiddleware sets up the middleware stack.
func assembleMiddleware() {
	http.Handle("/", JSONMiddleware(router))
}

func assembleRoutes() {
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
