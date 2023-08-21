package main

import (
	"log"
	"net/http"

	routers "github.com/thakkarnetram/go-server1/routes"
)

func main() {
	// router
	r := routers.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}