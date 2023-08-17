package main

import (
	"log"
	"net/http"

	"github.com/thakkarnetram/go-server1/helpers"
	routers "github.com/thakkarnetram/go-server1/routes"
)

func main() {
	// helpers
	helpers.Init()
	// router
	r := routers.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}