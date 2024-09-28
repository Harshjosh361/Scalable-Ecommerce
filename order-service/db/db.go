package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Database {
	clientOption := options.Client().ApplyURI("mongodb://mongo:27019/")
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	fmt.Println("successfully connected to database")

	return client.Database("Testing")
}
