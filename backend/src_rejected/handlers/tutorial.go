package handlers

import (
	"encoding/json"
	"time"

	"github.com/bugsfounder/bugsfounderweb/db/db_handler"
	"github.com/bugsfounder/bugsfounderweb/db/models"
	"go.mongodb.org/mongo-driver/bson"
)

// create
func CreateTutorial(tutorial models.Tutorial) error {
	tutorial.CreatedAt = time.Now()
	tutorial.UpdatedAt = time.Now()

	tutorialData, err := json.Marshal(tutorial)
	if err != nil {
		return err
	}

	var tutorialBson bson.M
	if err := json.Unmarshal(tutorialData, &tutorialBson); err != nil {
		return err
	}
	return db_handler.CreateUser(tutorialBson)
}

// read
func GetAllTutorials() ([]byte, error) {
	tutorials, err := db_handler.GetAllTutorials()
	if err != nil {
		return nil, err
	}

	return json.Marshal(tutorials)
}

func GetTutorialById(id string) ([]byte, error) {
	tutorial, err := db_handler.GetTutorialById(id)
	if err != nil {
		return nil, err
	}

	return json.Marshal(tutorial)
}

// update
func UpdateTutorial(id string, update models.Tutorial) ([]byte, error) {
	update.UpdatedAt = time.Now()

	updateData, err := json.Marshal(update)
	if err != nil {
		return nil, err
	}

	var updateBson bson.M
	if err := json.Unmarshal(updateData, &updateBson); err != nil {
		return nil, err
	}

	if err := db_handler.UpdateTutorial(id, updateBson); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "Tutorial updated successfully"})
}

// delete
func DeleteTutorial(id string) ([]byte, error) {
	if err := db_handler.DeleteTutorial(id); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "Tutorial deleted successfully"})
}

// hide
func HideTutorial(id string) ([]byte, error) {
	if err := db_handler.HideTutorial(id); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "Tutorial hidden successfully"})
}

// unhide
func UnHideTutorial(id string) ([]byte, error) {
	if err := db_handler.UnHideTutorial(id); err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{"message": "Tutorial unhidden successfully"})
}
