package handlers

import (
	"encoding/json"
	"time"

	"github.com/bugsfounder/bugsfounderweb/db/db_handler"
	"github.com/bugsfounder/bugsfounderweb/db/models"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// create
func CreateBlog(blog models.Blog) error {
	blogData, err := json.Marshal(blog)
	if err != nil {
		return err
	}

	var blogBson bson.M
	if err := json.Unmarshal(blogData, &blogBson); err != nil {
		return err
	}

	return db_handler.CreateBlog(blogBson)
}

// read
func GetAllBlogs() ([]byte, error) {
	blogs, err := db_handler.GetAllBlogs()
	if err != nil {
		return nil, err
	}
	logger.Logging().Info("blogs: %v", blogs)
	return json.Marshal(blogs)
}

func GetBlogById(id string) ([]byte, error) {
	blog, err := db_handler.GetBlogById(id)
	if err != nil {
		return nil, err
	}

	return json.Marshal(blog)
}

// update
func UpdateBlog(id string, update models.Blog) ([]byte, error) {
	update.UpdatedAt = time.Now()
	updateBson, err := bson.Marshal(update)
	if err != nil {
		return nil, err
	}
	if err := db_handler.UpdateBlog(id, updateBson); err != nil {
		return nil, err
	}
	return json.Marshal(map[string]string{"message": "Blog updated successfully"})
}

// delete
func DeleteBlog(id string) ([]byte, error) {
	if err := db_handler.DeleteBlog(id); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "Blog deleted successfully"})
}

func HideBlog(id string) ([]byte, error) {
	if err := db_handler.HideBlog(id); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "Blog hidden successfully"})
}
func UnHideBlog(id string) ([]byte, error) {
	if err := db_handler.UnHideBlog(id); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "Blog unhidden successfully"})
}
