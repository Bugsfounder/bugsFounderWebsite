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

func (client *Client) CreateOneBlog(blog *models.Blog) (*mongo.InsertOneResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")
	urlCount, err := client.SearchBlogURL(blog.Url)
	if err != nil {
		return nil, err
	}
	if urlCount > 0 {
		return nil, errors.New("blog url already exists")
	}
	// insert the blog into the collection
	result, err := collection.InsertOne(ctx, blog)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetAllBlogs(offset, limit int) ([]models.Blog, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	findOptions := options.Find()
	findOptions.SetSkip(int64(offset))
	findOptions.SetLimit(int64(limit))

	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		LOG.Error("%v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

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

func (client *Client) UpdateOneBlogByURL(blogURL string, updatedBlog *models.Blog) (*mongo.UpdateResult, *models.Blog, error) {
	LOG.Debug("")
	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	// set the updatedAt file
	updatedBlog.UpdatedAt = time.Now()

	// define the filter and update
	filter := bson.M{"url": blogURL}

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
		return nil, nil, err
	}

	return result, updatedBlog, nil
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
