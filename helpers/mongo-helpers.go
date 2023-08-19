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

var (
	url      string
	dbName   = "playlists"
	colName1 = "users"
	colName2 = "userPlaylist"
)

var collection1 *mongo.Collection
var collection2 *mongo.Collection

func Init() {
	// loading env
	err:=godotenv.Load("./.env") 
	if err != nil {  
		panic(err)
	}
	// setting url
	url = os.Getenv("MONGO_URI")
	// fmt.Println("URI ", url)
	// client options
	clientOption:=options.Client().ApplyURI(url)
	// connection
	client,err := mongo.Connect(context.TODO(),clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success")
	// db & collection 
	db := client.Database(dbName)
	collection1 = db.Collection(colName1)
	collection2 = db.Collection(colName2)
}	

func InsertUser(user *model.User) error  { 
	_,err := collection1.InsertOne(context.Background(),user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User inserted ", user)
	return nil
}

func FindUser (email string ) (bool,error) {
	// again filer
	filter:=bson.M{"email":email}
	// count 
	count,err := collection1.CountDocuments(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	return count>0 ,err
}

func EmailExists(email string ) (bool,error){
	// filter to find email
	filter:=bson.M{"email":email}
	// perform action
	count,err:=collection1.CountDocuments(context.Background(),filter);
	if err != nil {
		log.Fatal(err)
	}
	//  return count if big than 0
	return count>0 , err
}

func UserNameExists(username string ) (bool,error) {
	// filter the names
	filter:=bson.M{"username":username}
	// ok,err
	count,err:= collection1.CountDocuments(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0 , err
}