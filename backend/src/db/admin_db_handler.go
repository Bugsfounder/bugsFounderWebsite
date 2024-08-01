package db

import (
	"context"
	"errors"
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

func (client *Client) CreateOneAdmin(admin *models.Admin) (*mongo.InsertOneResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("admins")

	// check if username or email already exists
	isPresentUsername, _ := collection.CountDocuments(ctx, bson.M{"username": admin.Username})
	isPresentEmail, _ := collection.CountDocuments(ctx, bson.M{"email": admin.Email})
	if isPresentEmail > 0 || isPresentUsername > 0 {
		return nil, errors.New("email or username already exists")
	}
	// Insert the admin into the collection
	result, err := collection.InsertOne(ctx, admin)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetAdminCount() (int, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("admins")

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	LOG.Debug("AdminCount: %v", int(count))
	return int(count), nil
}

func (client *Client) DeleteAdminByUsername(username string) (*mongo.DeleteResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("admins")

	result, err := collection.DeleteOne(ctx, bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	return result, nil
}
