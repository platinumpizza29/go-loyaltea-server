package handlers

import (
	"loyaltea-server/internal/models"
	"loyaltea-server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	favService *services.FavoriteService
}

func NewFavoriteHandler(favService *services.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{
		favService: favService,
	}
}

// AddFavorite adds a favorite to the user's list of favorites.
func (h *FavoriteHandler) CreateFav(c *gin.Context) {
	var favorite *models.Favorite
	if err := c.ShouldBindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.favService.AddFavorite(favorite.UserID, favorite.OfferId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add favorite"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Favorite added successfully"})
}

// RemoveFavorite removes a favorite from the user's list of favorites.
func (h *FavoriteHandler) DeleteFav(c *gin.Context) {
	id := c.Param("id")
	if err := h.favService.RemoveUserFavorite(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove favorite"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Favorite removed successfully"})
}

// GetFavorites returns the list of favorites for a user.
func (h *FavoriteHandler) GetFav(c *gin.Context) {
	userID := c.Param("id")
	favorites, err := h.favService.GetFavorites(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get favorites"})
		return
	}
	c.JSON(http.StatusOK, favorites)
}

func (h *FavoriteHandler) UpdateFav(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Update favorite not implemented"})
}

// ClearFavorites clears all favorites for a user.
func (h *FavoriteHandler) ClearUserFavorites(c *gin.Context) {
	id := c.Param("id")
	if err := h.favService.RemoveUserFavorite(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear favorites"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Favorites cleared successfully"})
}
