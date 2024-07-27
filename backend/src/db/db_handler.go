package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Client_Obj *mongo.Client
}

func withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func (client *Client) DemoFunc() {
	ctx, cancel := withTimeout()
	defer cancel()
	// Access a specific collection
	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	doc := map[string]string{"title": "example", "content": "document"}

	// insert the document
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		LOG.Error("Failed to insert document: %v", err)
		return
	}
	LOG.Info("Document inserted successfully")
}

func (client *Client) CreateOneBlog() {

}
func (client *Client) GetAllBlogs() {

}
func (client *Client) GetOneBlogByID() {

}
func (client *Client) UpdateOneBlogById() {

}
func (client *Client) DeleteOneBlogById() {

}
func (client *Client) CreateOneTutorail() {

}
func (client *Client) GetAllTutorial() {

}
func (client *Client) GetOneTutorialByID() {

}
func (client *Client) UpdateOneTutorialById() {

}
func (client *Client) DeleteOneTutorialById() {

}
func (client *Client) GetAllUsers() {

}
func (client *Client) GetOneUserByUsername() {

}
func (client *Client) GetOneUserByEmail() {

}
func (client *Client) CreateOneUser() {

}
func (client *Client) UpdateOneUserByUsernameOrEmail() {

}
func (client *Client) DeleteOneUserByUsernameOrEmail() {

}
