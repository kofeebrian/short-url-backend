package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func DBConnect() *mongo.Client {
	opts := options.Client().ApplyURI("mongodb://db:27017")
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("Connection Failed to Database: ", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Connection Failed to Database: ", err)
	}

	log.Println("Connected to Database")

	return client
}
