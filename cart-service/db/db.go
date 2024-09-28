package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Database {
	clientOption := options.Client().ApplyURI("mongodb://mongo:27020/")
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("failed to ping database")
	}
	fmt.Println("Successfully connected to database")
	return client.Database("Testing")
}
