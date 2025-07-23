package handlers

import (
	"loyaltea-server/internal/models"
	"loyaltea-server/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	shopService *services.ShopService
}

func NewShopHandler(shopService *services.ShopService) *ShopHandler {
	return &ShopHandler{shopService: shopService}
}

// GET /shops/nearby?lat=<lat>&lng=<lng>&radius=<radius>
func (h *ShopHandler) GetNearbyShops(ctx *gin.Context) {
	latStr := ctx.Query("lat")
	lngStr := ctx.Query("lng")
	radiusStr := ctx.DefaultQuery("radius", "1000") // Default 1km

	if latStr == "" || lngStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "latitude and longitude are required"})
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}

	radius, err := strconv.Atoi(radiusStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid radius"})
		return
	}

	shops, err := h.shopService.FindNearbyShops(ctx, lng, lat, radius)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"shops": shops,
		"count": len(shops),
	})
}

// GET /shops?category=<category>&search=<query>
func (h *ShopHandler) GetShops(ctx *gin.Context) {
	category := ctx.Query("category")
	search := ctx.Query("search")

	var shops []*models.Shop
	var err error

	if search != "" {
		shops, err = h.shopService.SearchShops(ctx, search)
	} else {
		shops, err = h.shopService.GetShops(ctx, category)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"shops": shops,
		"count": len(shops),
	})
}

// GET /shops/:id
func (h *ShopHandler) GetShopByID(ctx *gin.Context) {
	shopID := ctx.Param("id")
	if shopID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Shop ID is required"})
		return
	}

	shop, err := h.shopService.GetShopByID(ctx, shopID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Shop not found"})
		return
	}

	ctx.JSON(http.StatusOK, shop)
}

// POST /shops (Admin only)
func (h *ShopHandler) CreateShop(ctx *gin.Context) {
	var shop models.Shop
	if err := ctx.ShouldBindJSON(&shop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate shop data
	if err := h.shopService.ValidateShopData(&shop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.shopService.CreateShop(ctx, &shop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, shop)
}

// PUT /shops/:id (Admin only)
func (h *ShopHandler) UpdateShop(ctx *gin.Context) {
	shopID := ctx.Param("id")
	if shopID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Shop ID is required"})
		return
	}

	var shop models.Shop
	if err := ctx.ShouldBindJSON(&shop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shop.ID = shopID

	// Validate shop data
	if err := h.shopService.ValidateShopData(&shop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.shopService.UpdateShop(ctx, &shop); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, shop)
}

// DELETE /shops/:id (Admin only)
func (h *ShopHandler) DeleteShop(ctx *gin.Context) {
	shopID := ctx.Param("id")
	if shopID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Shop ID is required"})
		return
	}

	if err := h.shopService.DeleteShop(ctx, shopID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Shop deleted successfully"})
}

// GET /shops/brand/:brand
func (h *ShopHandler) GetShopsByBrand(ctx *gin.Context) {
	brand := ctx.Param("brand")
	if brand == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Brand is required"})
		return
	}

	shops, err := h.shopService.GetShopsByBrand(ctx, brand)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"shops": shops,
		"count": len(shops),
		"brand": brand,
	})
}

// GET /shops/categories
func (h *ShopHandler) GetCategories(ctx *gin.Context) {
	// This would typically come from a database or configuration
	// For now, return common categories
	categories := []string{
		"coffee",
		"clothing",
		"grocery",
		"restaurant",
		"electronics",
		"pharmacy",
		"bookstore",
		"fitness",
		"beauty",
		"automotive",
	}

	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
