package services

import (
	"context"
	"loyaltea-server/internal/db"
	"loyaltea-server/internal/models"
)

type PlannerService struct {
	plannerDb *db.PlannerStruct
}

func NewPlannerService(plannerDb *db.PlannerStruct) *PlannerService {
	return &PlannerService{plannerDb: plannerDb}
}

func (p *PlannerService) GetAllGeoPoints(ctx context.Context) ([]*models.GeoPoint, error) {
	return p.plannerDb.GetAllGeoPoints(ctx)
}

// get stops by user id
func (p *PlannerService) GetStopsByUserID(ctx context.Context, userID string) ([]*models.Stop, error) {
	return p.plannerDb.ShoppingPlan(ctx, userID)
}

// create a new stop
func (p *PlannerService) CreateStop(ctx context.Context, stop *models.Stop) error {
	return p.plannerDb.CreatePlanner(ctx, stop)
}

// update a stop
func (p *PlannerService) UpdateStop(ctx context.Context, stop *models.Stop) error {
	return p.plannerDb.UpdateStop(ctx, stop)
}

// delete a stop
func (p *PlannerService) DeleteStop(ctx context.Context, stop *models.Stop) error {
	return p.plannerDb.DeleteStop(ctx, stop)
}
