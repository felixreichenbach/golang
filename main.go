package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// General Info: https://github.com/mongodb/mongo-go-driver/releases/tag/v1.0.0-rc1

// Test Branch

// Trainer will be used later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)

	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")

	// Insert documents

	ash := Trainer{"Ash", 10, "Pallet Town"}
	//misty := Trainer{"Misty", 10, "Cerulean City"}
	//brock := Trainer{"Brock", 15, "Pewter City"}

	insertResult, err := collection.InsertOne(context.TODO(), ash)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Disconnect from MongoDB

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
