package handlers

import (
	"encoding/json"

	"github.com/bugsfounder/bugsfounderweb/db/db_handler"
	"github.com/bugsfounder/bugsfounderweb/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// create
func CreateBlog(blogJSON []byte) ([]byte, error) {
	var blog bson.M
	if err := json.Unmarshal(blogJSON, &blog); err != nil {
		return nil, err
	}

	if err := db_handler.CreateBlog(blog); err != nil {
		return nil, err
	}
	return json.Marshal(map[string]string{"message": "blog created successfully"})
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
func UpdateBlog(id string, updateJSON []byte) ([]byte, error) {
	var update bson.M
	if err := json.Unmarshal(updateJSON, &update); err != nil {
		return nil, err
	}

	if err := db_handler.UpdateBlog(id, update); err != nil {
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
