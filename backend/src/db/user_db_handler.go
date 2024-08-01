package db

import (
	"context"
	"errors"
	"time"

	"github.com/bugsfounder/bugsfounderweb/models"
	"github.com/bugsfounder/bugsfounderweb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

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

func (client *Client) GetOneUserByUsernameOrEmail(usernameOrEmail, password string) (*models.User, error) {
	LOG.Debug("")

	ctx, cancel := withTimeout()
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("users")

	var user *models.User
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

	// compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// password do not match
		return nil, errors.New("invalid username/email or password")
	}

	// // Generate a JWT token
	// token, err := utils.GenerateJWT(user.Username, user.Email)
	// if err != nil {
	// 	return "", err
	// }

	return user, nil
}

func (client *Client) CreateOneUser(user *models.User) (*mongo.InsertOneResult, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("users")

	// check if email or username already exists
	emailFilter := bson.M{"email": user.Email}
	usernameFilter := bson.M{"username": user.Username}

	emailCount, err := collection.CountDocuments(ctx, emailFilter)
	if err != nil {
		return nil, err
	}
	if emailCount > 0 {
		return nil, errors.New("email already exists")
	}

	usernameCount, err := collection.CountDocuments(ctx, usernameFilter)
	if err != nil {
		return nil, err
	}
	if usernameCount > 0 {
		return nil, errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// set the createdAt and updated at field
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

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
