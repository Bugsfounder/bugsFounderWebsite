package db

import (
	"context"
	"time"

	"github.com/bugsfounder/bugsfounderweb/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (client *Client) BlacklistToken(token string, expiresAt time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blacklisted_tokens")

	blacklistedToken := models.BlacklistedToken{
		Token:     token,
		ExpiresAt: expiresAt,
	}

	_, err := collection.InsertOne(ctx, blacklistedToken)
	return err
}

func (client *Client) IsTokenBlacklisted(token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blacklisted_tokens")

	filter := bson.M{"token": token}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
