package services

import (
	"context"
	"loyaltea-server/internal/db"
	"loyaltea-server/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShopService struct {
	shopDb *db.ShopModel
}

func NewShopService(shopDb *db.ShopModel) *ShopService {
	return &ShopService{shopDb: shopDb}
}

func (s *ShopService) FindNearbyShops(ctx context.Context, longitude, latitude float64, maxDistance int) ([]*models.Shop, error) {
	return s.shopDb.FindNearbyShops(ctx, longitude, latitude, maxDistance)
}

func (s *ShopService) GetShops(ctx context.Context, category string) ([]*models.Shop, error) {
	return s.shopDb.GetShops(ctx, category)
}

func (s *ShopService) CreateShop(ctx context.Context, shop *models.Shop) error {
	now := time.Now()
	shop.ID = primitive.NewObjectID().Hex()
	shop.CreatedAt = now
	shop.UpdatedAt = now
	shop.IsActive = true

	// Validate required fields
	if shop.Name == "" {
		return ErrShopNameRequired
	}
	if shop.Location.Type == "" {
		shop.Location.Type = "Point"
	}
	if len(shop.Location.Coordinates) != 2 {
		return ErrInvalidLocation
	}
	if shop.PointsValue <= 0 {
		shop.PointsValue = 10 // Default points value
	}

	return s.shopDb.CreateShop(ctx, shop)
}

func (s *ShopService) GetShopByID(ctx context.Context, id string) (*models.Shop, error) {
	return s.shopDb.GetShopByID(ctx, id)
}

func (s *ShopService) UpdateShop(ctx context.Context, shop *models.Shop) error {
	shop.UpdatedAt = time.Now()
	return s.shopDb.UpdateShop(ctx, shop)
}

func (s *ShopService) DeleteShop(ctx context.Context, id string) error {
	return s.shopDb.DeleteShop(ctx, id)
}

func (s *ShopService) GetShopsByBrand(ctx context.Context, brand string) ([]*models.Shop, error) {
	return s.shopDb.GetShopsByBrand(ctx, brand)
}

func (s *ShopService) SearchShops(ctx context.Context, query string) ([]*models.Shop, error) {
	return s.shopDb.SearchShops(ctx, query)
}

func (s *ShopService) GetShopsByIDs(ctx context.Context, shopIDs []string) ([]models.Shop, error) {
	return s.shopDb.GetShopsByIDs(ctx, shopIDs)
}

// ValidateShopData validates shop data before creation/update
func (s *ShopService) ValidateShopData(shop *models.Shop) error {
	if shop.Name == "" {
		return ErrShopNameRequired
	}
	if shop.Address == "" {
		return ErrShopAddressRequired
	}
	if len(shop.Location.Coordinates) != 2 {
		return ErrInvalidLocation
	}
	if shop.Location.Coordinates[0] < -180 || shop.Location.Coordinates[0] > 180 {
		return ErrInvalidLongitude
	}
	if shop.Location.Coordinates[1] < -90 || shop.Location.Coordinates[1] > 90 {
		return ErrInvalidLatitude
	}
	return nil
}
