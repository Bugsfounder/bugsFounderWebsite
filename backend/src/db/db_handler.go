package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Client_Obj *mongo.Client
	Ctx        context.Context
}

func (client *Client) DemoFunc() {
	// Access a specific collection
	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	doc := map[string]string{"title": "example", "content": "document"}

	// insert the document
	_, err := collection.InsertOne(client.Ctx, doc)
	if err != nil {
		LOG.Error("Failed to insert document: %v", err)
		return
	}
	LOG.Info("Document inserted successfully")
}
