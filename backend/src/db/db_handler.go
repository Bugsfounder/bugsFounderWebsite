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
