package main

import (
	"context"
	"time"

	"github.com/bugsfounder/bugsfounderweb/logger"
	"github.com/bugsfounder/bugsfounderweb/router"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var LOG = logger.Logging()

func main() {
	LOG.Debug("Main of server")

	uri := "mongodb://bugsfounder:kubari@localhost:27017/"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		LOG.Error("Failed to create Client: %v", err)
	}

	// Create a context with a timeout to ensure the connection does not hang indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		LOG.Error("Failed to ping MongoDB: %v", err)
	}

	LOG.Info("Successfully connected and pinged MongoDB")

	// Access a specific collection
	collection := client.Database("bugsfounderDB").Collection("blogs")

	doc := map[string]string{"title": "example", "content": "document"}
	_, err = collection.InsertOne(ctx, doc)
	if err != nil {
		LOG.Error("Failed to insert document: %v", err)
	}
	LOG.Info("Document inserted successfully")

	server := gin.Default()

	router.ApiRoutes(server)
	server.Run(":8080")
}
