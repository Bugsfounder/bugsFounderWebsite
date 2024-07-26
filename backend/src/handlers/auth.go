package handlers

import (
	"encoding/json"
	"errors"

	"github.com/bugsfounder/bugsfounderweb/db/db_handler"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// create
func CreateUser(userJSON []byte) ([]byte, error) {
	var user bson.M
	if err := json.Unmarshal(userJSON, &user); err != nil {
		return nil, err
	}
	if err := db_handler.CreateUser(user); err != nil {
		return nil, err
	}
	return json.Marshal(map[string]string{"message": "user created successfully"})
}

// read
func GetUserByUsername(username string) ([]byte, error) {
	user, err := db_handler.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return json.Marshal(user)
}

func GetUserByEmail(email string) ([]byte, error) {
	user, err := db_handler.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return json.Marshal(user)
}

// update
func UpdateUser(id string, updateJSON []byte) ([]byte, error) {
	var update bson.M
	if err := json.Unmarshal(updateJSON, &update); err != nil {
		return nil, err
	}

	if err := db_handler.UpdateUser(id, update); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "User updated successfully"})
}

// delete
func DeleteUser(id string) ([]byte, error) {
	if err := db_handler.DeleteUser(id); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "User deleted successfully"})
}

// Signup
func Signup(userJSON []byte) ([]byte, error) {
	var user bson.M
	if err := json.Unmarshal(userJSON, &user); err != nil {
		return nil, err
	}

	password, ok := user["password"].(string)
	if !ok {
		return nil, errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user["password"] = string(hashedPassword)

	if err := db_handler.CreateUser(user); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "User signed up successfully"})
}

// Login
func Login(credentialsJSON []byte) ([]byte, error) {
	var credentials map[string]string
	if err := json.Unmarshal(credentialsJSON, &credentials); err != nil {
		return nil, err
	}

	username, ok := credentials["username"]
	if !ok {
		return nil, errors.New("username is required")
	}

	password, ok := credentials["password"]
	if !ok {
		return nil, errors.New("password is required")
	}

	user, err := db_handler.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	storedPassword, ok := user["password"].(string)
	if !ok {
		return nil, errors.New("stored password is invalid")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return json.Marshal(map[string]string{"message": "User logged in successfully"})
}

// Logout (for simplicity, just a placeholder)
func Logout() ([]byte, error) {
	return json.Marshal(map[string]string{"message": "User logged out successfully"})
}
