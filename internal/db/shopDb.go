package db

import (
	"context"
	"loyaltea-server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ShopModel struct {
	collection *mongo.Collection
}

func NewShopModel(db *mongo.Database) *ShopModel {
	return &ShopModel{
		collection: db.Collection("shops"),
	}
}

// Find nearby shops using MongoDB geospatial queries
func (s *ShopModel) FindNearbyShops(ctx context.Context, longitude, latitude float64, maxDistance int) ([]*models.Shop, error) {
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longitude, latitude},
				},
				"$maxDistance": maxDistance, // in meters
			},
		},
		"is_active": true,
	}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shops []*models.Shop
	if err := cursor.All(ctx, &shops); err != nil {
		return nil, err
	}

	return shops, nil
}

// Get all shops with optional filtering
func (s *ShopModel) GetShops(ctx context.Context, category string) ([]*models.Shop, error) {
	filter := bson.M{"is_active": true}
	if category != "" {
		filter["category"] = category
	}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shops []*models.Shop
	if err := cursor.All(ctx, &shops); err != nil {
		return nil, err
	}

	return shops, nil
}

func (s *ShopModel) CreateShop(ctx context.Context, shop *models.Shop) error {
	_, err := s.collection.InsertOne(ctx, shop)
	return err
}

func (s *ShopModel) GetShopByID(ctx context.Context, id string) (*models.Shop, error) {
	var shop models.Shop
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&shop)
	return &shop, err
}

func (s *ShopModel) UpdateShop(ctx context.Context, shop *models.Shop) error {
	_, err := s.collection.UpdateOne(ctx, bson.M{"_id": shop.ID}, bson.M{"$set": shop})
	return err
}

func (s *ShopModel) DeleteShop(ctx context.Context, id string) error {
	_, err := s.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"is_active": false}})
	return err
}

// Get shops by brand
func (s *ShopModel) GetShopsByBrand(ctx context.Context, brand string) ([]*models.Shop, error) {
	filter := bson.M{"brand": brand, "is_active": true}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shops []*models.Shop
	if err := cursor.All(ctx, &shops); err != nil {
		return nil, err
	}

	return shops, nil
}

// Search shops by name or brand
func (s *ShopModel) SearchShops(ctx context.Context, query string) ([]*models.Shop, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"is_active": true},
			{
				"$or": []bson.M{
					{"name": bson.M{"$regex": query, "$options": "i"}},
					{"brand": bson.M{"$regex": query, "$options": "i"}},
				},
			},
		},
	}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shops []*models.Shop
	if err := cursor.All(ctx, &shops); err != nil {
		return nil, err
	}

	return shops, nil
}

func (s *ShopModel) GetShopsByIDs(ctx context.Context, shopIDs []string) ([]models.Shop, error) {
	shopCollection := s.collection

	filter := bson.M{"_id": bson.M{"$in": shopIDs}} // no conversion needed
	cursor, err := shopCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shops []models.Shop
	if err := cursor.All(ctx, &shops); err != nil {
		return nil, err
	}

	if shops == nil {
		shops = []models.Shop{}
	}

	return shops, nil
}
