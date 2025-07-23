package db

import (
	"context"
	"loyaltea-server/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ShoppingPlanModel struct {
	collection *mongo.Collection
}

func NewShoppingPlanModel(db *mongo.Database) *ShoppingPlanModel {
	return &ShoppingPlanModel{
		collection: db.Collection("shopping_plans"),
	}
}

func (sp *ShoppingPlanModel) CreatePlan(ctx context.Context, plan *models.ShoppingPlan) error {
	_, err := sp.collection.InsertOne(ctx, plan)
	return err
}

func (sp *ShoppingPlanModel) GetPlanByID(ctx context.Context, id string) (*models.ShoppingPlan, error) {
	var plan models.ShoppingPlan
	err := sp.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&plan)
	return &plan, err
}

func (sp *ShoppingPlanModel) GetUserPlans(ctx context.Context, userID string) ([]*models.ShoppingPlan, error) {
	cursor, err := sp.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var plans []*models.ShoppingPlan
	if err := cursor.All(ctx, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}

func (sp *ShoppingPlanModel) UpdatePlan(ctx context.Context, plan *models.ShoppingPlan) error {
	plan.UpdatedAt = time.Now()
	_, err := sp.collection.UpdateOne(ctx, bson.M{"_id": plan.ID}, bson.M{"$set": plan})
	return err
}

func (sp *ShoppingPlanModel) DeletePlan(ctx context.Context, id string) error {
	_, err := sp.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// Mark a specific shop as visited in a plan
func (sp *ShoppingPlanModel) MarkShopVisited(ctx context.Context, planID, shopID string, pointsEarned int) error {
	visitedAt := time.Now()
	filter := bson.M{
		"_id":           planID,
		"stops.shop_id": shopID,
	}

	update := bson.M{
		"$set": bson.M{
			"stops.$.is_visited":    true,
			"stops.$.visited_at":    visitedAt,
			"stops.$.points_earned": pointsEarned,
			"updated_at":            visitedAt,
		},
	}

	_, err := sp.collection.UpdateOne(ctx, filter, update)
	return err
}

// Check if all stops in a plan are completed
func (sp *ShoppingPlanModel) CheckAndMarkPlanCompleted(ctx context.Context, planID string) error {
	plan, err := sp.GetPlanByID(ctx, planID)
	if err != nil {
		return err
	}

	allVisited := true
	for _, stop := range plan.Stops {
		if !stop.IsVisited {
			allVisited = false
			break
		}
	}

	if allVisited && !plan.IsCompleted {
		_, err := sp.collection.UpdateOne(
			ctx,
			bson.M{"_id": planID},
			bson.M{
				"$set": bson.M{
					"is_completed": true,
					"updated_at":   time.Now(),
				},
			},
		)
		return err
	}

	return nil
}

// Get completed plans for a user
func (sp *ShoppingPlanModel) GetCompletedPlans(ctx context.Context, userID string) ([]*models.ShoppingPlan, error) {
	cursor, err := sp.collection.Find(ctx, bson.M{
		"user_id":      userID,
		"is_completed": true,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var plans []*models.ShoppingPlan
	if err := cursor.All(ctx, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}

// Get active (non-completed) plans for a user
func (sp *ShoppingPlanModel) GetActivePlans(ctx context.Context, userID string) ([]*models.ShoppingPlan, error) {
	cursor, err := sp.collection.Find(ctx, bson.M{
		"user_id":      userID,
		"is_completed": false,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var plans []*models.ShoppingPlan
	if err := cursor.All(ctx, &plans); err != nil {
		return nil, err
	}

	return plans, nil
}
