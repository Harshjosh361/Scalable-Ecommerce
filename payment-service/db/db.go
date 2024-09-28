package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27021/")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// ping db
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("failed to ping database")
	}

	fmt.Print("Successfully connected to database")
	return client.Database("Testing")
}
