package main

import (
	"fmt"

	"github.com/bugsfounder/bugsfounderweb/db"
	"github.com/bugsfounder/bugsfounderweb/db/db_handler"
	"github.com/bugsfounder/bugsfounderweb/db/models"
)

func TextDB() {
	// Database credentials
	creds := db.NewDB_Credentials("bugsfounder", "kubari")

	// Connect to MongoDB
	client, err := db.ConnectToDB(creds)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	// Ensure connection is closed on exit
	defer db.DisconnectToDB(client)

	// DBHandler for blogs
	blogHandler := db_handler.NewDBHandler(client, "blogs")

	// Create a new blog
	newBlog := models.Blog{
		BlogID:  1,
		Title:   "First Blog",
		Content: "This is the content of the first blog.",
		Author:  "Author1",
		Tags:    []string{"go", "mongodb", "tutorial"},
	}
	err = blogHandler.CreateBlog(newBlog)
	if err != nil {
		fmt.Println("Error creating blog:", err)
	}

	// Read a blog
	readBlog, err := blogHandler.ReadBlog(1)
	if err != nil {
		fmt.Println("Error reading blog:", err)
	} else {
		fmt.Printf("Read Blog: %+v\n", readBlog)
	}

	// Update a blog
	updatedBlog := models.Blog{
		Title:   "Updated Blog Title",
		Content: "This is the updated content.",
		Author:  "Author1",
		Tags:    []string{"updated", "tutorial"},
	}
	err = blogHandler.UpdateBlog(1, updatedBlog)
	if err != nil {
		fmt.Println("Error updating blog:", err)
	}

	// Delete a blog
	err = blogHandler.DeleteBlog(1)
	if err != nil {
		fmt.Println("Error deleting blog:", err)
	}

	// DBHandler for users
	userHandler := db_handler.NewDBHandler(client, "users")

	// Create a new user
	newUser := models.User{
		ID:       1,
		Username: "user1",
		Email:    "user1@example.com",
		Password: "password1",
	}
	err = userHandler.CreateUser(newUser)
	if err != nil {
		fmt.Println("Error creating user:", err)
	}

	// Read a user
	readUser, err := userHandler.ReadUser(1)
	if err != nil {
		fmt.Println("Error reading user:", err)
	} else {
		fmt.Printf("Read User: %+v\n", readUser)
	}

	// Update a user
	updatedUser := models.User{
		Username: "updated_user1",
		Email:    "updated_user1@example.com",
	}
	err = userHandler.UpdateUser(1, updatedUser)
	if err != nil {
		fmt.Println("Error updating user:", err)
	}

	// Delete a user
	err = userHandler.DeleteUser(1)
	if err != nil {
		fmt.Println("Error deleting user:", err)
	}
}
