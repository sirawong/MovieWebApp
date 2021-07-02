package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// connect to database
func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://graphql123:321graphql@cluster0.ekieb.mongodb.net"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	// assign value in struct db
	DB.Client = client
	DB.Database = client.Database("sample_mflix")
}
