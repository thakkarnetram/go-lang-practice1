package routers

import (
	"github.com/gorilla/mux"
	usercontroller "github.com/thakkarnetram/go-server1/controllers"
)

func Router() *mux.Router {
	// router init 
	r:=mux.NewRouter()
	// def routers
	r.HandleFunc("/api/v1/signup",usercontroller.RegisterUser).Methods("POST")
	return r
}