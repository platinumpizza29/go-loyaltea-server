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

// GetNearbyShops godoc
// @Summary Get nearby shops
// @Description Retrieve shops within a radius of the given latitude and longitude
// @Tags Shops
// @Produce json
// @Param lat query number true "Latitude"
// @Param lng query number true "Longitude"
// @Param radius query int false "Search radius in meters (default: 1000)"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shops/nearby [get]
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

// GetShops godoc
// @Summary Get shops
// @Description Get all shops or search by category or query string
// @Tags Shops
// @Produce json
// @Param category query string false "Shop category"
// @Param search query string false "Search query"
// @Success 200 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shops [get]
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

// GetShopByID godoc
// @Summary Get shop by ID
// @Description Get a single shop by its ID
// @Tags Shops
// @Produce json
// @Param id path string true "Shop ID"
// @Success 200 {object} models.Shop
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /shops/{id} [get]
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

// CreateShop godoc
// @Summary Create a new shop (admin)
// @Description Create a new shop entry. Admin access only.
// @Tags Shops
// @Accept json
// @Produce json
// @Param shop body models.Shop true "Shop data"
// @Success 201 {object} models.Shop
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shops [post]
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

// UpdateShop godoc
// @Summary Update shop (admin)
// @Description Update an existing shop by ID. Admin access only.
// @Tags Shops
// @Accept json
// @Produce json
// @Param id path string true "Shop ID"
// @Param shop body models.Shop true "Updated shop data"
// @Success 200 {object} models.Shop
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shops/{id} [put]
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

// DeleteShop godoc
// @Summary Delete a shop (admin)
// @Description Delete a shop by ID. Admin access only.
// @Tags Shops
// @Produce json
// @Param id path string true "Shop ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shops/{id} [delete]
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

// GetShopsByBrand godoc
// @Summary Get shops by brand
// @Description Get all shops that belong to a specific brand
// @Tags Shops
// @Produce json
// @Param brand path string true "Brand name"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /shops/brand/{brand} [get]
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

// GetCategories godoc
// @Summary Get shop categories
// @Description Returns a list of predefined shop categories
// @Tags Shops
// @Produce json
// @Success 200 {object} gin.H
// @Router /shops/categories [get]
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

type ShopIDsRequest struct {
	IDs []string `json:"ids" binding:"required"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

// GetShopsBatchHandler godoc
// @Summary      Get shops by a list of IDs
// @Description  Fetch a batch of shop documents based on provided comma-separated shop IDs
// @Tags         Shops
// @Accept       json
// @Produce      json
// @Param        body  body      ShopIDsRequest  true  "List of shop IDs"
// @Success      200   {array}   models.Shop
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /shops/batch [post]
func (h *ShopHandler) GetShopsBatchHandler(ctx *gin.Context) {
	var req ShopIDsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload, expected JSON with 'ids'"})
		return
	}

	shops, err := h.shopService.GetShopsByIDs(ctx.Request.Context(), req.IDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch shops"})
		return
	}

	ctx.JSON(http.StatusOK, shops)
}
