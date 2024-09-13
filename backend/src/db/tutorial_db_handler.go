package db

import (
	"context"
	"errors"
	"time"

	"github.com/bugsfounder/bugsfounderweb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (client *Client) CreateOneTutorial(tutorial *models.Tutorial) (*mongo.InsertOneResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	tutorialCount, err := client.SearchTutorialURL(tutorial.Url)
	if err != nil {
		return nil, err
	}
	if tutorialCount > 0 {
		return nil, errors.New("tutorial url already exists")
	}
	// set the createdAt and updated at field
	tutorial.CreatedAt = time.Now()
	tutorial.UpdatedAt = time.Now()
	tutorial.Sub_Tutorials = []models.Sub_Tutorial{}

	// insert the blog into the collection
	result, err := collection.InsertOne(ctx, tutorial)
	if err != nil {
		return nil, err
	}
	return result, nil

}

/*
	func (client *Client) GetAllBlogs(offset, limit int) ([]models.Blog, error) {
		var blogs []models.Blog
		for cursor.Next(ctx) {
			var blog models.Blog
			if err = cursor.Decode(&blog); err != nil {
				LOG.Error("%v", err)
				return nil, err
			}
			blogs = append(blogs, blog)
		}

		if err := cursor.Err(); err != nil {
			LOG.Error("%v", err)
			return nil, err
		}

		return blogs, nil
	}
*/

func (client *Client) GetAllTutorial(offset, limit int) ([]models.Tutorial, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()
	// select the database and collection
	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	findOptions := options.Find()
	findOptions.SetSkip(int64(offset))
	findOptions.SetLimit(int64(limit))

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
		if err = cursor.Decode(&tutorial); err != nil {
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

func (client *Client) UpdateOneTutorialByURL(tutorialURL string, updatedTutorial *models.Tutorial) (*mongo.UpdateResult, *models.Tutorial, error) {
	LOG.Debug("")

	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// set the updatedAt file
	updatedTutorial.UpdatedAt = time.Now()

	// define the filter and update
	filter := bson.M{"url": tutorialURL}

	updateFields := bson.M{}
	if updatedTutorial.Title != "" {
		updateFields["title"] = updatedTutorial.Title
	}
	if updatedTutorial.Description != "" {
		updateFields["content"] = updatedTutorial.Description
	}
	if updatedTutorial.Author != "" {
		updateFields["author"] = updatedTutorial.Author
	}
	if updatedTutorial.Url != "" {
		updateFields["url"] = updatedTutorial.Url
	}
	if len(updatedTutorial.Sub_Tutorials) > 0 {
		updateFields["sub_tutorials"] = updatedTutorial.Sub_Tutorials
	}
	if len(updatedTutorial.Tags) > 0 {
		updateFields["tags"] = updatedTutorial.Tags
	}
	updateFields["updated_at"] = updatedTutorial.UpdatedAt

	update := bson.M{
		"$set": updateFields,
	}
	// update the blog in the collection
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, nil, err
	}

	return result, updatedTutorial, nil
}

func (client *Client) DeleteOneTutorialByURL(tutorialURL string) (*mongo.DeleteResult, error) {
	LOG.Debug("")

	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// define the filter
	filter := bson.M{"url": tutorialURL}

	// delete the tutorial from the collection
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return result, err

}

// SUB-TUTORIAL
func (client *Client) CreateSubTutorial(tutorialURL string, subTutorial *models.Sub_Tutorial) (*mongo.UpdateResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	subTutorialCount, err := client.SearchSubTutorialURL(tutorialURL, subTutorial.Url)
	if err != nil {
		return nil, err
	}
	if subTutorialCount > 0 {
		return nil, errors.New("sub tutorial url already exits")
	}

	// Set the CreatedAt and UpdatedAt fields for the sub-tutorial
	subTutorial.CreatedAt = time.Now()
	subTutorial.UpdatedAt = time.Now()

	// Define the filter to find the tutorial by URL
	filter := bson.M{"url": tutorialURL}

	// Define the update to push the new sub-tutorial into the sub_tutorials array
	update := bson.M{
		"$push": bson.M{
			"sub_tutorials": subTutorial,
		},
	}

	// Update the tutorial document
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) GetSubTutorialByURL(tutorialURL, subTutorialURL string) (*models.Sub_Tutorial, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// define the filter to match the tutorial and the sub-tutorial
	filter := bson.M{
		"url":               tutorialURL,
		"sub_tutorials.url": subTutorialURL,
	}

	// Find the tutorial document
	var tutorial models.Tutorial
	err := collection.FindOne(ctx, filter).Decode(&tutorial)
	if err != nil {
		return nil, err
	}

	// Extract the sub-tutorial
	for _, subTutorial := range tutorial.Sub_Tutorials {
		if subTutorial.Url == subTutorialURL {
			return &subTutorial, nil
		}
	}

	return nil, mongo.ErrNoDocuments

}

func (client *Client) UpdateSubTutorialByURL(tutorialURL, subTutorialURL string, updatedSubTutorial *models.Sub_Tutorial) (*mongo.UpdateResult, error) {
	LOG.Debug("")

	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// set the updatedAt field
	updatedSubTutorial.UpdatedAt = time.Now()

	// define the filter to match the tutorial and sub-tutorial
	filter := bson.M{
		"url":               tutorialURL,
		"sub_tutorials.url": subTutorialURL,
	}

	// prepare the update fields
	updateFields := bson.M{}
	if updatedSubTutorial.Title != "" {
		updateFields["sub_tutorials.$.title"] = updatedSubTutorial.Title
	}
	if updatedSubTutorial.Content != "" {
		updateFields["sub_tutorials.$.content"] = updatedSubTutorial.Content
	}
	if updatedSubTutorial.Author != "" {
		updateFields["sub_tutorials.$.author"] = updatedSubTutorial.Author
	}
	if updatedSubTutorial.Url != "" {
		updateFields["sub_tutorials.$.url"] = updatedSubTutorial.Url
	}
	if len(updatedSubTutorial.Tags) > 0 {
		updateFields["sub_tutorials.$.tags"] = updatedSubTutorial.Tags
	}
	updateFields["sub_tutorials.$.updated_at"] = updatedSubTutorial.UpdatedAt

	update := bson.M{
		"$set": updateFields,
	}

	// update the sub-tutorial in the collection
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) DeleteSubTutorialByURL(tutorialURL, subTutorialURL string) (*mongo.UpdateResult, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// define the filter and update
	filter := bson.M{"url": tutorialURL}
	update := bson.M{
		"$pull": bson.M{"sub_tutorials": bson.M{"url": subTutorialURL}},
	}

	// update the tutorial in the collection
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}
