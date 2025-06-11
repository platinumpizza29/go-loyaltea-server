package services

import (
	"context"
	"loyaltea-server/internal/models"
)

type OfferService struct {
	offerModel *models.OfferModel
}

func NewOfferService(offerModel *models.OfferModel) *OfferService {
	return &OfferService{
		offerModel: offerModel,
	}
}

func (s *OfferService) CreateOffer(ctx context.Context, offer *models.Offer) error {
	return s.offerModel.Create(ctx, offer)
}
