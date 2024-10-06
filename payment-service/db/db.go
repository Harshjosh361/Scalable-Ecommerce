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
	clientOptions := options.Client().ApplyURI(mongoURI)

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
