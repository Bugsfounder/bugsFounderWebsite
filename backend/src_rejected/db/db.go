package db

import (
	"context"
	"fmt"

	"github.com/bugsfounder/bugsfounderweb/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var LOG = logger.Logging()

type DB_Credentials struct {
	Username string
	Password string
}

func NewDB_Credentials(username string, password string) *DB_Credentials {
	return &DB_Credentials{
		Username: username,
		Password: password,
	}
}

func ConnectToDB(db_credentials *DB_Credentials) (*mongo.Client, error) {
	username := db_credentials.Username
	password := db_credentials.Password
	uri := fmt.Sprintf("mongodb://%v:%v@localhost:27017", username, password)
	clientOptions := options.Client().ApplyURI(uri)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		LOG.Error("%v", err)
		return nil, err
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		LOG.Error("%v", err)
		return nil, err
	}
	LOG.Info("Connected to MongoDB!")
	return client, nil
}

func DisconnectToDB(client *mongo.Client) {
	// Ensure connection is closed
	if err := client.Disconnect(context.TODO()); err != nil {
		LOG.Error("%v", err)
	}
}

func CreateDatabaseAndCollections(client *mongo.Client, databaseName string) error {
	db := client.Database(databaseName)

	// Create collections
	collections := []string{"blogs", "tutorials", "users", "tutorials"}
	for _, collectionName := range collections {
		err := db.CreateCollection(context.TODO(), collectionName)
		if err != nil {
			return fmt.Errorf("error creating collection %s: %v", collectionName, err)
		}
	}

	return nil
}
