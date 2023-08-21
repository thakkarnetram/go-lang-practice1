package usercontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/thakkarnetram/go-server1/helpers"
	"github.com/thakkarnetram/go-server1/model"
	"golang.org/x/crypto/bcrypt"
)

var db *helpers.DatabaseHelper
var err error

func init() {
	db,err = helpers.OpenDatabaseHelper()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected " , db)
}

func VerifyPass(userPassword string , responsePassword string ) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(responsePassword),[]byte(userPassword))
	if err != nil {
		return errors.New("Password is incorrect  ")
	}
	return nil
}	

func RegisterUser (res http.ResponseWriter , req *http.Request) {
	//  headers
	res.Header().Set("Content-Type","application/json")
	// model
	var user model.User
	// if user input empty
	if req.ContentLength == 0 {
		json.NewEncoder(res).Encode("Cannot send empty data")
		return
	}
	_=json.NewDecoder(req.Body).Decode(&user)
	if user.IsEmpty() {
		json.NewEncoder(res).Encode("All fields are required to Sign up ")
		return
	}
	// if any one field empty 
	if user.IsEmailPresent() {
		json.NewEncoder(res).Encode("Email should be provided")
		return
	}
	if user.IsNamePresent() {
		json.NewEncoder(res).Encode("User Name should be provided")
		return
	}
	if user.IsPassPresent(){
		json.NewEncoder(res).Encode("Password should be provided")
		return
	}
	// if email exists
	exists,err:=db.EmailExists(user.Email)
	if err != nil {
		http.Error(res,"Failed to check email exists or not ", http.StatusInternalServerError)
		return
	}
	if exists {
		json.NewEncoder(res).Encode("Email exists please login")
		return
	}
	// if username exists 
	exist,err:=db.UserNameExists(user.Name)
	if err != nil {
		http.Error(res,"Failed to check username in db " , http.StatusInternalServerError)
		return 
	}
	if exist {
		json.NewEncoder(res).Encode("Username already taken")
		return
	}
	json.NewDecoder(req.Body).Decode(&user) // send reference of it 
	user.Password = getHash([]byte(user.Password))
	err = db.InsertUser(&user)
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

func LoginUser (res  http.ResponseWriter , req *http.Request) {
	res.Header().Set("Content-Type","application/json")
	// declare var 
	var user *model.User
	// decode body 
	_ = json.NewDecoder(req.Body).Decode(&user)
	// validations 
	if  req.ContentLength == 0 {
		json.NewEncoder(res).Encode("Cannot Send empty data ")
		return
	}
	if user.IsEmpty() {
		json.NewEncoder(res).Encode("Email and Password required")
		return
	}
	if user.IsEmailPresent(){
		json.NewEncoder(res).Encode("Email field cannot be empty")
		return
	}
	if user.IsPassPresent(){
		json.NewEncoder(res).Encode("Password field cannot be empty")
		return
	}
	// if user exists 
	exists , err := db.EmailExists(user.Email)
	if err != nil {
		http.Error(res,"Failed to Connect to the server" , http.StatusInternalServerError)
		return
	}
	// if email not there 
	if exists == false {
		json.NewEncoder(res).Encode("Email doesnt exist , Please Create an account ")
	}
	// if yes 
	if exists {
		// get the user 
		foundUser , err := db.GetUserByEmail(user.Email)
		if err != nil {
			http.Error(res,"Server error finding the user " , http.StatusInternalServerError)
			return
		}

		// compare password 
		err = VerifyPass(user.Password , foundUser.Password)
		if err != nil {
			response:=struct {
				Message string `json:"message"`
			} {
				Message: err.Error(),
			}
			json.NewEncoder(res).Encode(response)
			return
		}
		// if pass valid 
		response := struct {
			Message string `json:"message"`
		} {
			Message: "Logged In ",
		}
		json.NewEncoder(res).Encode(response)
	}
}

func getHash(pwd []byte) string {
	hash,err:= bcrypt.GenerateFromPassword(pwd,bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}