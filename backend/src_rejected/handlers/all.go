package handlers

// import (
// 	"encoding/json"

// 	"github.com/bugsfounder/bugsfounderweb/db/db_handler"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// // create
// func CreateUser(userJSON []byte) ([]byte, error) {
// 	var user bson.M
// 	if err := json.Unmarshal(userJSON, &user); err != nil {
// 		return nil, err
// 	}
// 	if err := db_handler.CreateUser(user); err != nil {
// 		return nil, err
// 	}
// 	return json.Marshal(map[string]string{"message": "user created successfully"})
// }

// // read
// func GetUserByUsername(username string) ([]byte, error) {
// 	user, err := db_handler.GetUserByUsername(username)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(user)
// }

// func GetUserByEmail(email string) ([]byte, error) {
// 	user, err := db_handler.GetUserByEmail(email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(user)
// }

// // update
// func UpdateUser(id string, updateJSON []byte) ([]byte, error) {
// 	var update bson.M
// 	if err := json.Unmarshal(updateJSON, &update); err != nil {
// 		return nil, err
// 	}

// 	if err := db_handler.UpdateUser(id, update); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "User updated successfully"})
// }

// // delete
// func DeleteUser(id string) ([]byte, error) {
// 	if err := db_handler.DeleteUser(id); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "User deleted successfully"})
// }
// func CreateBlog(blogJSON []byte) ([]byte, error) {
// 	var blog bson.M
// 	if err := json.Unmarshal(blogJSON, &blog); err != nil {
// 		return nil, err
// 	}

// 	if err := db_handler.CreateBlog(blog); err != nil {
// 		return nil, err
// 	}
// 	return json.Marshal(map[string]string{"message": "blog created successfully"})
// }

// // read
// func GetAllBlogs() ([]byte, error) {
// 	blogs, err := db_handler.GetAllBlogs()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(blogs)
// }

// func GetBlogById(id string) ([]byte, error) {
// 	blog, err := db_handler.GetBlogById(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(blog)
// }

// // update
// func UpdateBlog(id string, updateJSON []byte) ([]byte, error) {
// 	var update bson.M
// 	if err := json.Unmarshal(updateJSON, &update); err != nil {
// 		return nil, err
// 	}

// 	if err := db_handler.UpdateBlog(id, update); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Blog updated successfully"})
// }

// // delete
// func DeleteBlog(id string) ([]byte, error) {
// 	if err := db_handler.DeleteBlog(id); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Blog deleted successfully"})
// }

// func HideBlog(id string) ([]byte, error) {
// 	if err := db_handler.HideBlog(id); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Blog hidden successfully"})
// }
// func UnHideBlog(id string) ([]byte, error) {
// 	if err := db_handler.UnHideBlog(id); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Blog unhidden successfully"})
// }

// // create
// func CreateTutorial(tutorialJSON []byte) ([]byte, error) {
// 	var tutorial bson.M
// 	if err := json.Unmarshal(tutorialJSON, &tutorial); err != nil {
// 		return nil, err
// 	}

// 	if err := db_handler.CreateTutorial(tutorial); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Tutorial created successfully"})
// }

// // read
// func GetAllTutorials() ([]byte, error) {
// 	tutorials, err := db_handler.GetAllTutorials()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(tutorials)
// }

// func GetTutorialById(id string) ([]byte, error) {
// 	tutorial, err := db_handler.GetTutorialById(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(tutorial)
// }

// // update

// func UpdateTutorial(id string, updateJSON []byte) ([]byte, error) {
// 	var update bson.M
// 	if err := json.Unmarshal(updateJSON, &update); err != nil {
// 		return nil, err
// 	}

// 	if err := db_handler.UpdateTutorial(id, update); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Tutorial updated successfully"})
// }

// // delete
// func DeleteTutorial(id string) ([]byte, error) {
// 	if err := db_handler.DeleteTutorial(id); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Tutorial deleted successfully"})
// }

// func HideTutorial(id string) ([]byte, error) {
// 	if err := db_handler.HideTutorial(id); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Tutorial hidden successfully"})
// }

// func UnHideTutorial(id string) ([]byte, error) {
// 	if err := db_handler.UnHideTutorial(id); err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(map[string]string{"message": "Tutorial unhidden successfully"})
// }
