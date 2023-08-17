package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/thakkarnetram/go-server1/helpers"
	"github.com/thakkarnetram/go-server1/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser (res http.ResponseWriter , req *http.Request) {
	//  headers
	res.Header().Set("Content-Type","application/json")
	// model
	var user model.User
	if req.ContentLength == 0 {
		json.NewEncoder(res).Encode("Cannot send empty data")
		return
	}
	if user.IsEmpty() {
		json.NewEncoder(res).Encode("All fields are required to Sign up ")
		return
	}
	json.NewDecoder(req.Body).Decode(&user) // send reference of it 
	user.Password = getHash([]byte(user.Password))
	err := helpers.InsertUser(&user)
	if err != nil {
		http.Error(res,"Failed to register ", http.StatusInternalServerError)
		return
	}
	response := struct {
		Message string `json:"message"`
	} {
		Message: "User register successfully ",
	}
	json.NewEncoder(res).Encode(response)
}

func getHash(pwd []byte) string {
	hash,err:= bcrypt.GenerateFromPassword(pwd,bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}