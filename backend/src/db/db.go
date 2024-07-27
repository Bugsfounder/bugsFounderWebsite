package db

import (
	"context"
	"time"

	"github.com/bugsfounder/bugsfounderweb/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var LOG = logger.Logging()

func ConnectToDBAndGetClientCtxCancelErr() (*mongo.Client, context.Context, context.CancelFunc, error) {
	uri := "mongodb://bugsfounder:kubari@localhost:27017/"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		LOG.Error("Failed to create Client: %v", err)
		return nil, nil, nil, err

	}

	// Create a context with a timeout to ensure the connection does not hang indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		LOG.Error("Failed to ping MongoDB: %v", err)
		cancel()
		return nil, nil, nil, err
	}

	LOG.Info("Successfully connected and pinged MongoDB")
	return client, ctx, cancel, err
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			LOG.Error("%v", err)
			panic(err)
		}
	}()
}

func Ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	LOG.Info("Connected Successfully")
	return nil
}
