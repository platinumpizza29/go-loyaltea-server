package services

import (
	"context"
	"loyaltea-server/internal/db"
	"loyaltea-server/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShoppingPlanService struct {
	planDb *db.ShoppingPlanModel
	userDb *db.UserModel
	shopDb *db.ShopModel
}

func NewShoppingPlanService(planDb *db.ShoppingPlanModel, userDb *db.UserModel, shopDb *db.ShopModel) *ShoppingPlanService {
	return &ShoppingPlanService{
		planDb: planDb,
		userDb: userDb,
		shopDb: shopDb,
	}
}

func (sp *ShoppingPlanService) CreatePlan(ctx context.Context, plan *models.ShoppingPlan) error {
	now := time.Now()
	plan.ID = primitive.NewObjectID().Hex()
	plan.CreatedAt = now
	plan.UpdatedAt = now
	plan.IsCompleted = false

	// Validate required fields
	if plan.Name == "" {
		return ErrPlanNameRequired
	}
	if plan.UserID == "" {
		return ErrUserIDRequired
	}

	// Validate that all shops exist
	for i, stop := range plan.Stops {
		shop, err := sp.shopDb.GetShopByID(ctx, stop.ShopID)
		if err != nil {
			return ErrShopNotFound
		}
		if !shop.IsActive {
			return ErrShopNotActive
		}
		// Initialize stop values
		plan.Stops[i].IsVisited = false
		plan.Stops[i].PointsEarned = 0
	}

	return sp.planDb.CreatePlan(ctx, plan)
}

func (sp *ShoppingPlanService) GetPlanByID(ctx context.Context, id string) (*models.ShoppingPlan, error) {
	return sp.planDb.GetPlanByID(ctx, id)
}

func (sp *ShoppingPlanService) GetUserPlans(ctx context.Context, userID string) ([]*models.ShoppingPlan, error) {
	return sp.planDb.GetUserPlans(ctx, userID)
}

func (sp *ShoppingPlanService) GetActivePlans(ctx context.Context, userID string) ([]*models.ShoppingPlan, error) {
	return sp.planDb.GetActivePlans(ctx, userID)
}

func (sp *ShoppingPlanService) GetCompletedPlans(ctx context.Context, userID string) ([]*models.ShoppingPlan, error) {
	return sp.planDb.GetCompletedPlans(ctx, userID)
}

func (sp *ShoppingPlanService) UpdatePlan(ctx context.Context, plan *models.ShoppingPlan) error {
	return sp.planDb.UpdatePlan(ctx, plan)
}

func (sp *ShoppingPlanService) DeletePlan(ctx context.Context, id string) error {
	return sp.planDb.DeletePlan(ctx, id)
}

// MarkShopVisited marks a shop as visited and awards points to the user
func (sp *ShoppingPlanService) MarkShopVisited(ctx context.Context, planID, userID, shopID string) error {
	// Get the shop to know how many points to award
	shop, err := sp.shopDb.GetShopByID(ctx, shopID)
	if err != nil {
		return err
	}

	// Get the plan to validate ownership and check if already visited
	plan, err := sp.planDb.GetPlanByID(ctx, planID)
	if err != nil {
		return err
	}

	// Verify the plan belongs to the user
	if plan.UserID != userID {
		return ErrUnauthorizedAccess
	}

	// Check if the shop is in the plan and not already visited
	shopFound := false
	alreadyVisited := false
	for _, stop := range plan.Stops {
		if stop.ShopID == shopID {
			shopFound = true
			if stop.IsVisited {
				alreadyVisited = true
			}
			break
		}
	}

	if !shopFound {
		return ErrShopNotInPlan
	}
	if alreadyVisited {
		return ErrShopAlreadyVisited
	}

	// Mark the shop as visited in the plan
	err = sp.planDb.MarkShopVisited(ctx, planID, shopID, shop.PointsValue)
	if err != nil {
		return err
	}

	// Award points to the user
	err = sp.userDb.AddPoints(ctx, userID, shop.PointsValue)
	if err != nil {
		return err
	}

	// Check if all stops are completed and mark plan as completed
	err = sp.planDb.CheckAndMarkPlanCompleted(ctx, planID)
	if err != nil {
		return err
	}

	return nil
}

// AddShopToPlan adds a new shop to an existing plan
func (sp *ShoppingPlanService) AddShopToPlan(ctx context.Context, planID, userID, shopID string) error {
	// Get the plan to validate ownership
	plan, err := sp.planDb.GetPlanByID(ctx, planID)
	if err != nil {
		return err
	}

	// Verify the plan belongs to the user
	if plan.UserID != userID {
		return ErrUnauthorizedAccess
	}

	// Check if plan is already completed
	if plan.IsCompleted {
		return ErrPlanAlreadyCompleted
	}

	// Validate that the shop exists and is active
	shop, err := sp.shopDb.GetShopByID(ctx, shopID)
	if err != nil {
		return ErrShopNotFound
	}
	if !shop.IsActive {
		return ErrShopNotActive
	}

	// Check if shop is already in the plan
	for _, stop := range plan.Stops {
		if stop.ShopID == shopID {
			return ErrShopAlreadyInPlan
		}
	}

	// Add the new stop to the plan
	newStop := models.PlanStop{
		ShopID:       shopID,
		IsVisited:    false,
		PointsEarned: 0,
	}
	plan.Stops = append(plan.Stops, newStop)

	return sp.planDb.UpdatePlan(ctx, plan)
}

// RemoveShopFromPlan removes a shop from an existing plan
func (sp *ShoppingPlanService) RemoveShopFromPlan(ctx context.Context, planID, userID, shopID string) error {
	// Get the plan to validate ownership
	plan, err := sp.planDb.GetPlanByID(ctx, planID)
	if err != nil {
		return err
	}

	// Verify the plan belongs to the user
	if plan.UserID != userID {
		return ErrUnauthorizedAccess
	}

	// Check if plan is already completed
	if plan.IsCompleted {
		return ErrPlanAlreadyCompleted
	}

	// Find and remove the shop from stops
	var newStops []models.PlanStop
	shopFound := false
	for _, stop := range plan.Stops {
		if stop.ShopID != shopID {
			newStops = append(newStops, stop)
		} else {
			shopFound = true
			// If the shop was already visited, we need to subtract the points
			if stop.IsVisited {
				err = sp.userDb.AddPoints(ctx, userID, -stop.PointsEarned)
				if err != nil {
					return err
				}
			}
		}
	}

	if !shopFound {
		return ErrShopNotInPlan
	}

	plan.Stops = newStops
	return sp.planDb.UpdatePlan(ctx, plan)
}

// GetPlanProgress returns the progress of a shopping plan
func (sp *ShoppingPlanService) GetPlanProgress(ctx context.Context, planID string) (*PlanProgress, error) {
	plan, err := sp.planDb.GetPlanByID(ctx, planID)
	if err != nil {
		return nil, err
	}

	totalStops := len(plan.Stops)
	visitedStops := 0
	totalPoints := 0

	for _, stop := range plan.Stops {
		if stop.IsVisited {
			visitedStops++
			totalPoints += stop.PointsEarned
		}
	}

	progress := &PlanProgress{
		PlanID:         plan.ID,
		TotalStops:     totalStops,
		VisitedStops:   visitedStops,
		TotalPoints:    totalPoints,
		IsCompleted:    plan.IsCompleted,
		CompletionRate: 0,
	}

	if totalStops > 0 {
		progress.CompletionRate = float64(visitedStops) / float64(totalStops) * 100
	}

	return progress, nil
}

// GetUserStats returns statistics for a user's shopping activities
func (sp *ShoppingPlanService) GetUserStats(ctx context.Context, userID string) (*UserStats, error) {
	allPlans, err := sp.GetUserPlans(ctx, userID)
	if err != nil {
		return nil, err
	}

	completedPlans, err := sp.GetCompletedPlans(ctx, userID)
	if err != nil {
		return nil, err
	}

	totalPoints, err := sp.userDb.GetUserPoints(ctx, userID)
	if err != nil {
		return nil, err
	}

	totalVisits := 0
	for _, plan := range allPlans {
		for _, stop := range plan.Stops {
			if stop.IsVisited {
				totalVisits++
			}
		}
	}

	stats := &UserStats{
		UserID:         userID,
		TotalPlans:     len(allPlans),
		CompletedPlans: len(completedPlans),
		TotalPoints:    totalPoints,
		TotalVisits:    totalVisits,
	}

	return stats, nil
}

// PlanProgress represents the progress of a shopping plan
type PlanProgress struct {
	PlanID         string  `json:"plan_id"`
	TotalStops     int     `json:"total_stops"`
	VisitedStops   int     `json:"visited_stops"`
	TotalPoints    int     `json:"total_points"`
	IsCompleted    bool    `json:"is_completed"`
	CompletionRate float64 `json:"completion_rate"`
}

// UserStats represents statistics for a user's shopping activities
type UserStats struct {
	UserID         string `json:"user_id"`
	TotalPlans     int    `json:"total_plans"`
	CompletedPlans int    `json:"completed_plans"`
	TotalPoints    int    `json:"total_points"`
	TotalVisits    int    `json:"total_visits"`
}
