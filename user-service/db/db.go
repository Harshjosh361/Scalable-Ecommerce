package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to mongodb:", err)
	}

	// checking server is reachable and functioning
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB", err)
	}

	fmt.Print("Successfuly connected to mongodb")
	return client.Database("Testing")
}
