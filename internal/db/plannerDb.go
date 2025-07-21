package db

import (
	"context"
	"loyaltea-server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PlannerStruct struct {
	collection *mongo.Collection
}

func NewPlannerStruct(db *mongo.Database) *PlannerStruct {
	return &PlannerStruct{
		collection: db.Collection("planner"),
	}
}

func (p *PlannerStruct) GetAllGeoPoints(ctx context.Context) ([]*models.GeoPoint, error) {
	var geoPoints []*models.GeoPoint
	cursor, err := p.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &geoPoints); err != nil {
		return nil, err
	}
	return geoPoints, nil
}

func (p *PlannerStruct) CreatePlanner(ctx context.Context, planner *models.Stop) error {
	_, err := p.collection.InsertOne(ctx, planner)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlannerStruct) ShoppingPlan(ctx context.Context, userID string) ([]*models.Stop, error) {
	var stops []*models.Stop
	cursor, err := p.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &stops); err != nil {
		return nil, err
	}
	return stops, nil
}

// update the stop
func (p *PlannerStruct) UpdateStop(ctx context.Context, stop *models.Stop) error {
	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": stop.ID}, bson.M{"$set": stop})
	if err != nil {
		return err
	}
	return nil
}

// delete the stop
func (p *PlannerStruct) DeleteStop(ctx context.Context, stop *models.Stop) error {
	_, err := p.collection.DeleteOne(ctx, bson.M{"_id": stop.ID})
	if err != nil {
		return err
	}
	return nil
}
