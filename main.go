package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page for instagram Clone.\n Routes below \n /createuser for creating new user \n /createpost for create new Post by user")
}

func createUser(w http.ResponseWriter, r *http.Request) {

	//connection to database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://priyankdb:priyank123@instagram.bwzus.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Fprintf(w, "This is userpage")
	instadatabase := client.Database("insta")
	userCollection := instadatabase.Collection("user")
	//postCollection := instadatabase.Collection("post")

	userResult, err := userCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Pragya"},
		{Key: "email", Value: "pragyalovespriyank@gmail.com"},
		{Key: "password", Value: "ilupriyank"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userResult.InsertedID)
	json.NewEncoder(w).Encode(userResult.InsertedID)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://priyankdb:priyank123@instagram.bwzus.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Fprintf(w, "This is userpage")
	instadatabase := client.Database("insta")
	userCollection := instadatabase.Collection("user")
	postCollection := instadatabase.Collection("post")

	userResult, err := userCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Priyank Gandu"},
		{Key: "email", Value: "priyankg@gmail.com"},
		{Key: "password", Value: "priyankggggg"},
	})

	postResult, err := postCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "post", Value: userResult.InsertedID},
			{Key: "caption", Value: "my lovely memory with pragya :("},
			{Key: "imageUrl", Value: "htttp://someimageurl"},
			{Key: "time", Value: primitive.Timestamp{}},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(postResult.InsertedIDs)
	json.NewEncoder(w).Encode(postResult.InsertedIDs)

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/createpost", createPost)
	http.HandleFunc("/createuser", createUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	handleRequests()

}
