package db

// create a new database connection for mongo db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
)

// ConnectDB establishes a connection to MongoDB
func ConnectDB(uri string, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return err
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return err
	}

	// Set global variables
	Client = client
	Database = client.Database(dbName)

	log.Println("Connected to MongoDB!")
	return nil
}

// CloseDB closes the MongoDB connection
func CloseDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
		return err
	}

	log.Println("Disconnected from MongoDB!")
	return nil
}
