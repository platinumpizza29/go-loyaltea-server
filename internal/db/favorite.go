package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FavModel struct {
	collection *mongo.Collection
}

func NewFavoriteModel(client *mongo.Database) *FavModel {
	return &FavModel{
		collection: client.Collection("favorites"),
	}
}

func (m *FavModel) AddFavorite(userID string, offerID string) error {
	_, err := m.collection.InsertOne(context.Background(), bson.M{
		"_id":           primitive.NewObjectID().Hex(),
		"user_id":       userID,
		"offer_id":      offerID,
		"offer_created": time.Now(),
	})
	return err
}

// get the user's favorite offers
func (m *FavModel) GetUserFavorites(userID string) ([]string, error) {
	cursor, err := m.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var favorites []string
	for cursor.Next(context.Background()) {
		var favorite bson.M
		if err := cursor.Decode(&favorite); err != nil {
			return nil, err
		}
		if offerID, ok := favorite["offer_id"].(string); ok {
			favorites = append(favorites, offerID)
		}
	}
	return favorites, nil
}

// remove a favorite offer
func (m *FavModel) RemoveFavorite(id string) error {
	_, err := m.collection.DeleteOne(context.Background(), bson.M{
		"_id": id,
	})
	return err
}

// check if a user has a favorite offer
func (m *FavModel) HasFavorite(userID string, offerID string) (bool, error) {
	count, err := m.collection.CountDocuments(context.Background(), bson.M{
		"user_id":  userID,
		"offer_id": offerID,
	})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
