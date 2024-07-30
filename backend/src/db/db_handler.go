package db

import (
	"context"
	"time"

	"github.com/bugsfounder/bugsfounderweb/models"
	"github.com/bugsfounder/bugsfounderweb/utils"
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

func (client *Client) CreateOneBlog(blog *models.Blog) (*mongo.InsertOneResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	// set the createdAt and updated at field
	// blog.CreatedAt = time.Now()
	// blog.UpdatedAt = time.Now()

	// insert the blog into the collection
	result, err := collection.InsertOne(ctx, blog)
	if err != nil {
		return nil, err
	}
	return result, nil
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

func (client *Client) UpdateOneBlogByURL(blogURL string, updatedBlog *models.Blog) (*mongo.UpdateResult, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	// set the updatedAt file
	updatedBlog.UpdatedAt = time.Now()

	// define the filter and update
	filter := bson.M{"url": blogURL}
	// update := bson.M{
	// 	"$set": bson.M{
	// 		"title":      updatedBlog.Title,
	// 		"content":    updatedBlog.Content,
	// 		"author":     updatedBlog.Author,
	// 		"tags":       updatedBlog.Tags,
	// 		"updated_at": updatedBlog.UpdatedAt,
	// 	},
	// }

	updateFields := bson.M{}
	if updatedBlog.Title != "" {
		updateFields["title"] = updatedBlog.Title
	}
	if updatedBlog.Content != "" {
		updateFields["content"] = updatedBlog.Content
	}
	if updatedBlog.Author != "" {
		updateFields["author"] = updatedBlog.Author
	}
	if updatedBlog.Url != "" {
		updateFields["url"] = updatedBlog.Url
	}
	if len(updatedBlog.Tags) > 0 {
		updateFields["tags"] = updatedBlog.Tags
	}
	updateFields["updated_at"] = updatedBlog.UpdatedAt

	update := bson.M{
		"$set": updateFields,
	}
	// update the blog in the collection
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) DeleteOneBlogByURL(blogURL string) (*mongo.DeleteResult, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	// define the filter
	filter := bson.M{"url": blogURL}

	// delete the blog from the collection
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) CreateOneTutorial(tutorial *models.Tutorial) (*mongo.InsertOneResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// set the createdAt and updated at field
	// tutorial.CreatedAt = time.Now()
	// tutorial.UpdatedAt = time.Now()

	// insert the blog into the collection
	result, err := collection.InsertOne(ctx, tutorial)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (client *Client) CreateSubTutorial(tutorialURL string, subTutorial *models.Sub_Tutorial) (*mongo.UpdateResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

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

func (client *Client) UpdateOneTutorialByURL(tutorialURL string, updatedTutorial *models.Tutorial) (*mongo.UpdateResult, error) {
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
		return nil, err
	}

	return result, nil
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

func (client *Client) GetOneUserByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	LOG.Debug("")

	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("users")

	var user models.User

	var filter bson.M

	if utils.IsValidEmail(usernameOrEmail) {
		filter = bson.M{"email": usernameOrEmail}
	} else {
		filter = bson.M{"username": usernameOrEmail}
	}

	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (client *Client) CreateOneUser(user *models.User) (*mongo.InsertOneResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("users")

	// set the createdAt and updated at field
	// user.CreatedAt = time.Now()
	// user.UpdatedAt = time.Now()

	// insert the blog into the collection
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (client *Client) UpdateOneUserByUsernameOrEmail(emailOrUsername string, updatedUser *models.User) (*mongo.UpdateResult, error) {
	LOG.Debug("")

	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("users")

	// set the updatedAt file
	updatedUser.UpdatedAt = time.Now()

	// define the filter and update
	var filter bson.M
	if utils.IsValidEmail(emailOrUsername) {
		filter = bson.M{"email": emailOrUsername}
	} else {
		filter = bson.M{"username": emailOrUsername}

	}
	// update := bson.M{
	// 	"$set": bson.M{
	// 		"email":      updatedUser.Email,
	// 		"username":   updatedUser.Username,
	// 		"password":   updatedUser.Password,
	// 		"updated_at": updatedUser.UpdatedAt,
	// 	},
	// }

	// build the update object based on non-empty fields
	updateFields := bson.M{}

	if updatedUser.Email != "" {
		updateFields["email"] = updatedUser.Email
	}
	if updatedUser.Username != "" {
		updateFields["username"] = updatedUser.Username
	}
	if updatedUser.Password != "" {
		updateFields["password"] = updatedUser.Password
	}

	update := bson.M{
		"$set": updateFields,
	}

	// update the blog in the collection
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) DeleteOneUserByUsernameOrEmail(usernameOrEmail string) (*mongo.DeleteResult, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("users")

	var filter bson.M

	if utils.IsValidEmail(usernameOrEmail) {
		filter = bson.M{"email": usernameOrEmail}
	} else {
		filter = bson.M{"username": usernameOrEmail}
	}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
