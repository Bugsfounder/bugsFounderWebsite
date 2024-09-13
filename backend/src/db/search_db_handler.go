package db

import (
	"context"
	"strings"
	"time"

	"github.com/bugsfounder/bugsfounderweb/models"
	"go.mongodb.org/mongo-driver/bson"
)

// search
func (client *Client) SearchBlogURL(url string) (int64, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("blogs")

	// check if email or username already exists
	BlogFilter := bson.M{"url": url}

	SearchBlogURLCount, err := collection.CountDocuments(ctx, BlogFilter)
	if err != nil {
		return -1, err
	}

	// insert the blog into the collection
	return SearchBlogURLCount, nil
}
func (client *Client) SearchTutorialURL(url string) (int64, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// check if email or username already exists
	tutorialFilter := bson.M{"url": url}

	tutorialURLCount, err := collection.CountDocuments(ctx, tutorialFilter)
	if err != nil {
		return -1, err
	}

	// insert the blog into the collection
	return tutorialURLCount, nil
}
func (client *Client) SearchSubTutorialURL(tutorialURL, subTutorialURL string) (int64, error) {
	LOG.Debug("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// check if email or username already exists
	tutorialFilter := bson.M{"url": tutorialURL, "sub_tutorials.url": subTutorialURL}

	tutorialURLCount, err := collection.CountDocuments(ctx, tutorialFilter)
	if err != nil {
		return -1, err
	}

	// insert the blog into the collection
	return tutorialURLCount, nil
}

func (client *Client) Search(query string) ([]interface{}, error) {
	queryWords := strings.Split(query, " ")

	var results []interface{}

	// search for each query words in blog
	blogs, err := client.searchCollection("blogs", queryWords)
	if err != nil {
		return nil, err
	}
	results = append(results, map[string]interface{}{"blogs": blogs})

	// search for each query words in tutorial
	tutorials, err := client.searchCollection("tutorials", queryWords)
	if err != nil {
		return nil, err
	}
	results = append(results, map[string]interface{}{"tutorials": tutorials})

	// search for each query words in subTutorial
	subTutorials, err := client.searchSubTutorials(queryWords)
	if err != nil {
		return nil, err
	}
	results = append(results, map[string]interface{}{"subTutorials": subTutorials})

	return results, nil

}

func (client *Client) searchCollection(collectionName string, queryWords []string) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection(collectionName)

	var filters []bson.M
	for _, word := range queryWords {
		filters = append(filters, bson.M{"$or": []bson.M{
			{"title": bson.M{"$regex": word, "$options": "i"}},
			{"content": bson.M{"$regex": word, "$options": "i"}},
		}})
	}
	filter := bson.M{"$and": filters}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []interface{}
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// searchSubTutorials searches sub-tutorials within tutorials
func (client *Client) searchSubTutorials(queryWords []string) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := client.Client_Obj.Database("bugsfounderDB").Collection("tutorials")

	// Create filter for tutorials
	var tutorialFilters []bson.M
	for _, word := range queryWords {
		tutorialFilters = append(tutorialFilters, bson.M{"$or": []bson.M{
			{"title": bson.M{"$regex": word, "$options": "i"}},
			{"description": bson.M{"$regex": word, "$options": "i"}},
		}})
	}
	tutorialFilter := bson.M{"$and": tutorialFilters}

	cursor, err := collection.Find(ctx, tutorialFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var subTutorials []interface{}
	for cursor.Next(ctx) {
		var tutorial models.Tutorial
		if err := cursor.Decode(&tutorial); err != nil {
			return nil, err
		}

		// Filter sub-tutorials that match the query
		var matchingSubTutorials []models.Sub_Tutorial
		for _, subTutorial := range tutorial.Sub_Tutorials {
			if containsAllWords(subTutorial.Title, queryWords) || containsAllWords(subTutorial.Content, queryWords) {
				matchingSubTutorials = append(matchingSubTutorials, subTutorial)
			}
		}

		if len(matchingSubTutorials) > 0 {
			subTutorials = append(subTutorials, map[string]interface{}{
				"tutorial":     tutorial.Title, // You can include more tutorial details if needed
				"subTutorials": matchingSubTutorials,
			})
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return subTutorials, nil
}

// containsAllWords checks if all queryWords are present in the text
func containsAllWords(text string, queryWords []string) bool {
	for _, word := range queryWords {
		if !strings.Contains(strings.ToLower(text), strings.ToLower(word)) {
			return false
		}
	}
	return true
}
