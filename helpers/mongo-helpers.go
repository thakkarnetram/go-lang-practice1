package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thakkarnetram/go-server1/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseHelper struct {
	usersCollection *mongo.Collection
	playlistCollection *mongo.Collection
}

func OpenDatabaseHelper() (*DatabaseHelper,error) {
	// env
	err:=godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
	// mongo connection 
	url := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(url)
	client , err := mongo.Connect(context.Background(),clientOptions)
	if err != nil {
		log.Fatal(err)
	}	

	// create collections 
	usersCollection := client.Database("playlists").Collection("users")
	playlistCollection:= client.Database("playlists").Collection("userPlaylists")

	// return 
	return &DatabaseHelper{
		usersCollection: usersCollection,
		playlistCollection: playlistCollection,
	},nil
}


// insert method 
func (db *DatabaseHelper ) InsertUser (user *model.User) error {
	_,err := db.usersCollection.InsertOne(context.Background(),user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user " , user)
	return nil
}

// email exists  method 
func (db *DatabaseHelper) EmailExists (email string) (bool,error) {
	filter:= bson.M{"email":email}
	count , err := db.usersCollection.CountDocuments(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	return count>0,err
}

// user name exists method 
func (db *DatabaseHelper) UserNameExists(username string ) (bool,error) {
	filter := bson.M{"username":username}
	count , err := db.usersCollection.CountDocuments(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0 , err
}

// getting user by email method
func (db *DatabaseHelper ) GetUserByEmail (email string ) (*model.User , error ) {
	filter := bson.M{"email":email}
	var user model.User
	err := db.usersCollection.FindOne(context.Background(),filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return &user,nil
}
