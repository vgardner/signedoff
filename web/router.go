package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/EconomistDigitalSolutions/ramlapi"
	"github.com/gorilla/mux"
)

var router *mux.Router

// NewRouter wires up the middleware and routes.
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

	for _, up := range ep.URIParameters {
		if up.Pattern != "" {
			path = strings.Replace(
				path,
				fmt.Sprintf("{%s}", up.Key),
				fmt.Sprintf("{%s:%s}", up.Key, up.Pattern),
				1)
		}
	}

	route := router.
		Methods(ep.Verb).
		Path(path).
		Handler(RouteMap[ep.Handler])

	for _, param := range ep.QueryParameters {
		if param.Required {
			if param.Pattern != "" {
				route.Queries(param.Key, fmt.Sprintf("{%s:%s}", param.Key, param.Pattern))
			} else {
				route.Queries(param.Key, "")
			}
		}
	}
}
