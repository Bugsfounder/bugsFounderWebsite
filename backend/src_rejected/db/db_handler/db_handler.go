package db_handler

import (
	"context"

	"github.com/bugsfounder/bugsfounderweb/db"
	"go.mongodb.org/mongo-driver/bson"
)

var database = "bugsfounderDB"
var username = "bugsfounder"
var password = "kubari"
var dbCre *db.DB_Credentials = db.NewDB_Credentials(username, password)
var client, err = db.ConnectToDB(dbCre)

// init
func init() {
	if err != nil {
		db.LOG.Error("%v", err)
	} else {
		err := db.CreateDatabaseAndCollections(client, "bugsFounderDB")
		if err != nil {
			db.LOG.Error("%v", err)
		}
	}
}

// create
func CreateBlog(blog bson.M) error {
	collection := client.Database(database).Collection("blogs")
	_, err := collection.InsertOne(context.TODO(), blog)
	return err
}

func CreateTutorial(tutorial bson.M) error {
	collection := client.Database(database).Collection("tutorials")
	_, err := collection.InsertOne(context.TODO(), tutorial)
	return err
}

func CreateUser(user bson.M) error {
	collection := client.Database(database).Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

// read
func GetAllBlogs() ([]bson.M, error) {
	collection := client.Database(database).Collection("blogs")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	return results, err
}
func GetAllTutorials() ([]bson.M, error) {
	collection := client.Database(database).Collection("tutorials")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	return results, err
}
func GetTutorialById(id string) (bson.M, error) {
	collection := client.Database(database).Collection("tutorials")
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	return result, err
}
func GetBlogById(id string) (bson.M, error) {
	collection := client.Database(database).Collection("blogs")
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	return result, err
}

func GetUserByUsername(username string) (bson.M, error) {
	collection := client.Database(database).Collection("users")
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)
	return result, err
}
func GetUserByEmail(email string) (bson.M, error) {
	collection := client.Database(database).Collection("users")
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"username": email}).Decode(&result)
	return result, err
}

// update
func UpdateBlog(id string, update bson.M) error {
	collection := client.Database(database).Collection("blogs")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func UpdateTutorial(id string, update bson.M) error {
	collection := client.Database(database).Collection("tutorials")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})
	return err
}
func UpdateUser(id string, update bson.M) error {
	collection := client.Database(database).Collection("users")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

// delete
func DeleteBlog(id string) error {
	collection := client.Database(database).Collection("blogs")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
func DeleteTutorial(id string) error {
	collection := client.Database(database).Collection("tutorials")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
func DeleteUser(id string) error {
	collection := client.Database(database).Collection("users")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func HideBlog(id string) error {
	collection := client.Database("bugsfounderDB").Collection("blogs")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"hidden": true}})
	return err
}

func HideTutorial(id string) error {
	collection := client.Database("bugsfounderDB").Collection("tutorials")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"hidden": true}})
	return err
}
func UnHideBlog(id string) error {
	collection := client.Database("bugsfounderDB").Collection("blogs")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"hidden": false}})
	return err
}

func UnHideTutorial(id string) error {
	collection := client.Database("bugsfounderDB").Collection("tutorials")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"hidden": false}})
	return err
}
