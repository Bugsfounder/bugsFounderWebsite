package db

import (
	"context"
	"errors"
	"time"

	"github.com/bugsfounder/bugsfounderweb/admin"
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

func (client *Client) CreateOneAdmin(adminModel *models.Admin) (*mongo.InsertOneResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("admins")

	// check if username or email already exists
	isPresentUsername, _ := collection.CountDocuments(ctx, bson.M{"username": adminModel.Username})
	isPresentEmail, _ := collection.CountDocuments(ctx, bson.M{"email": adminModel.Email})
	if isPresentEmail > 0 || isPresentUsername > 0 {
		return nil, errors.New("email or username already exists")
	}

	// Check admin count
	counts, err := client.GetAdminCount()
	if err != nil {
		return nil, err
	}

	// Validate and add admin mode
	err = admin.ValidateAndAddAdminMode(adminModel.AdminMode, counts) // Pass the correct AdminMode value
	if err != nil {
		return nil, err
	}

	// Insert the admin into the collection
	result, err := collection.InsertOne(ctx, adminModel)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) GetAdminCount() (admin.AdminCount, error) { // Return admin.AdminCount
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("admins")

	var counts admin.AdminCount
	var err error
	ManagerCount, err := collection.CountDocuments(ctx, bson.M{"adminmode": admin.Manager})
	if err != nil {
		return counts, err
	}

	EditorsCount, err := collection.CountDocuments(ctx, bson.M{"adminmode": admin.Editor})
	if err != nil {
		return counts, err
	}

	ViewersCount, err := collection.CountDocuments(ctx, bson.M{"adminmode": admin.Viewer})
	if err != nil {
		return counts, err
	}

	counts.Managers = int(ManagerCount)
	counts.Editors = int(EditorsCount)
	counts.Viewers = int(ViewersCount)
	counts.Total = counts.Managers + counts.Editors + counts.Viewers
	return counts, nil
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
