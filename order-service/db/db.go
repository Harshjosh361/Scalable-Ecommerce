package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() *mongo.Database {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("failed to get env")
	}
	clientOption := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	fmt.Println("successfully connected to database")

	return client.Database("Testing")
}
