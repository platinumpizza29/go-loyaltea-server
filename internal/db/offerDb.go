package db

import (
	"context"
	"loyaltea-server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type OfferModel struct {
	collection *mongo.Collection
}

// NewOfferModel creates a new OfferModel instance
func NewOfferModel(db *mongo.Database) *OfferModel {
	return &OfferModel{
		collection: db.Collection("offers"),
	}
}

func (m *OfferModel) GetAll(ctx context.Context) ([]*models.Offer, error) {
	filter := bson.M{}
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var offers []*models.Offer
	if err := cursor.All(ctx, &offers); err != nil {
		return nil, err
	}
	return offers, nil
}

func (m *OfferModel) GetByID(ctx context.Context, id string) (*models.Offer, error) {
	var offer models.Offer
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&offer)
	if err != nil {
		return nil, err
	}
	return &offer, nil
}
