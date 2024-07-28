package db

import (
	"context"
	"time"

	"github.com/bugsfounder/bugsfounderweb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Client_Obj *mongo.Client
}

func withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func (client *Client) DemoFunc() {
	LOG.Debug("")
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
	LOG.Debug("")

}
func (client *Client) GetAllBlogs() ([]models.Blog, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	// select the database and collection
	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		LOG.Error("%v", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	// read all blogs
	var allBlogsFromDatabase []models.Blog
	for cursor.Next(ctx) {
		var blog models.Blog
		err := cursor.Decode(&blog)
		if err != nil {
			LOG.Error("Failed to decode document: %v", err)
			return nil, err
		}

		allBlogsFromDatabase = append(allBlogsFromDatabase, blog)
	}

	if err := cursor.Err(); err != nil {
		LOG.Error("Cursor error: %v", err)
		return nil, err
	}

	return allBlogsFromDatabase, nil

}
func (client *Client) GetOneBlogByURL(blogURL string) (*models.Blog, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	var blog models.Blog

	err := collection.FindOne(ctx, bson.M{"url": blogURL}).Decode(&blog)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &blog, nil
}
func (client *Client) UpdateOneBlogByURL() {
	LOG.Debug("")

}
func (client *Client) DeleteOneBlogByURL() {
	LOG.Debug("")

}
func (client *Client) CreateOneTutorail() {
	LOG.Debug("")

}
func (client *Client) GetAllTutorial() ([]models.Tutorial, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()
	// select the database and collection
	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		LOG.Error("%v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// read all blogs
	var allTutorialsFromDatabase []models.Tutorial
	for cursor.Next(ctx) {
		var tutorial models.Tutorial
		err := cursor.Decode(&tutorial)
		if err != nil {
			LOG.Error("Failed to decode document: %v", err)
			return nil, err
		}

		allTutorialsFromDatabase = append(allTutorialsFromDatabase, tutorial)
	}

	if err := cursor.Err(); err != nil {
		LOG.Error("Cursor error: %v", err)
		return nil, err
	}

	return allTutorialsFromDatabase, nil

}
func (client *Client) GetOneTutorialByURL(tutorial_url string) (*models.Tutorial, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	var tutorial models.Tutorial

	err := collection.FindOne(ctx, bson.M{"url": tutorial_url}).Decode(&tutorial)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &tutorial, nil
}
func (client *Client) handleGetOneTutorialBySubURL() {
	LOG.Debug("")
}
func (client *Client) UpdateOneTutorialByURL() {
	LOG.Debug("")

}
func (client *Client) DeleteOneTutorialByURL() {
	LOG.Debug("")

}
func (client *Client) GetAllUsers() ([]models.User, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()
	// select the database and collection
	collection := client.Client_Obj.Database("bugsfounderDB").Collection("users")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		LOG.Error("%v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// read all blogs
	var allUsersFromDatabase []models.User
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			LOG.Error("Failed to decode document: %v", err)
			return nil, err
		}

		allUsersFromDatabase = append(allUsersFromDatabase, user)
	}

	if err := cursor.Err(); err != nil {
		LOG.Error("Cursor error: %v", err)
		return nil, err
	}

	return allUsersFromDatabase, nil
}
func (client *Client) GetOneUserByUsername() {
	LOG.Debug("")

}
func (client *Client) GetOneUserByEmail() {
	LOG.Debug("")

}
func (client *Client) CreateOneUser() {
	LOG.Debug("")

}
func (client *Client) UpdateOneUserByUsernameOrEmail() {
	LOG.Debug("")

}
func (client *Client) DeleteOneUserByUsernameOrEmail() {
	LOG.Debug("")

}
