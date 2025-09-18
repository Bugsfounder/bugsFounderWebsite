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
	// Replace <db_password> with your actual password
	uri := "mongodb+srv://bugsfounder2021_db_user:OkdnRjEQ50RBQkNV@cluster0.vi2nkfp.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		LOG.Error("Failed to create Client: %v", err)
		return nil, nil, nil, err
	}

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Ping
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		LOG.Error("Failed to ping MongoDB: %v", err)
		cancel()
		return nil, nil, nil, err
	}

	LOG.Info("Successfully connected and pinged MongoDB Atlas ✅")
	return client, ctx, cancel, nil
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
	LOG.Info("Connected Successfully to MongoDB Atlas ✅")
	return nil
}
