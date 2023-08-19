package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string  		  `json:"username,omitempty" bson:"username,omitempty"`
	Email string 		  `json:"email,omitempty" bson:"email,omitempty"` 	
	Password string 	  `json:"password,omitempty" bson:"password,omitempty"` 	
}

func (u *User) IsEmpty() bool {
	return u.Name == "" && u.Email == "" && u.Password == ""
}

func (u *User) IsEmailPresent() bool {
	return u.Email == ""
}

func (u *User) IsNamePresent() bool {
	return u.Name == ""
}

func (u *User) IsPassPresent() bool {
	return u.Password == ""
}
