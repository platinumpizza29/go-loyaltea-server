package services

import (
	"context"
	"loyaltea-server/internal/db"
	"loyaltea-server/internal/models"
)

type OfferService struct {
	offerModel *db.OfferModel
}

func NewOfferService(offerModel *db.OfferModel) *OfferService {
	return &OfferService{
		offerModel: offerModel,
	}
}

func (s *OfferService) GetOffers(ctx context.Context) ([]*models.Offer, error) {
	return s.offerModel.GetAll(ctx)
}

// GetOfferByID handles GET requests and returns an offer by ID
func (s *OfferService) GetOfferByID(ctx context.Context, id string) (*models.Offer, error) {
	return s.offerModel.GetByID(ctx, id)
}
