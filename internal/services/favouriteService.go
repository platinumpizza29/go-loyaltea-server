package services

import (
	"loyaltea-server/internal/db"
)

type FavoriteService struct {
	favModel *db.FavModel
}

func NewFavoriteService(favModel *db.FavModel) *FavoriteService {
	return &FavoriteService{
		favModel: favModel,
	}
}

func (s *FavoriteService) AddFavorite(userID string, offerID string) error {
	return s.favModel.AddFavorite(userID, offerID)
}

func (s *FavoriteService) RemoveUserFavorite(Id string) error {
	return s.favModel.RemoveFavorite(Id)
}

func (s *FavoriteService) HasFavorite(userID string, offerID string) (bool, error) {
	return s.favModel.HasFavorite(userID, offerID)
}

func (s *FavoriteService) GetFavorites(userID string) ([]string, error) {
	return s.favModel.GetUserFavorites(userID)
}
