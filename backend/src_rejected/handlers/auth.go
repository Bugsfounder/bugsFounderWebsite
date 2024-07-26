package handlers

import (
	"encoding/json"
	"errors"

	"github.com/bugsfounder/bugsfounderweb/db/db_handler"
	"github.com/bugsfounder/bugsfounderweb/db/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return db_handler.CreateUser(userData)
}

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

func AuthenticateUser(credentials map[string]string) ([]byte, error) {
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
