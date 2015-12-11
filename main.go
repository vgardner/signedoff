package main

import (
	"log"
	"net/http"

	"github.com/vgardner/signedoff-api/web"
)

func main() {
	web.NewRouter()
	log.Println("running...")
	log.Fatal(http.ListenAndServe(":3002", nil))
}
