package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	database mongo.Database
	client   mongo.Client
}

func CreateRepository(mongodbURI string) MongoRepository {
	clientOptions := options.Client().ApplyURI(mongodbURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	database := client.Database("logs")
	return MongoRepository{
		database: *database,
		client:   *client,
	}

}

func (repo *MongoRepository) Insert(logLine LogLine) {
	collection := repo.database.Collection("attacks")

	// Define the filter to find the document
	filter := bson.M{"source": logLine.Source}

	// Create an update document with the values you want to set
	update := bson.M{
		"$set": bson.M{
			"source":        logLine.Source,
			"sourceDetails": logLine.SourceDetails,
			"lastUpdate":    logLine.Time,
		},
		"$push": bson.M{
			"logEvents": logLine,
		},
	}

	opts := options.Update().SetUpsert(true)

	// Perform the upsert operation
	result, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Print the number of documents modified (0 for insert, 1 for update)
	fmt.Printf("Upserted %v document(s)\n", result.ModifiedCount)
}

func (repo *MongoRepository) Disconnect() {
	if err := repo.client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}
