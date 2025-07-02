package handlers

import (
	"log"
	"loyaltea-server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type OfferHandler struct {
	offerService *services.OfferService
}

func NewOfferHandler(offerService *services.OfferService) *OfferHandler {
	return &OfferHandler{
		offerService: offerService,
	}
}

// GetOffers handles GET requests and returns all offers
func (h *OfferHandler) GetOffers(c *gin.Context) {
	offers, err := h.offerService.GetOffers(c.Request.Context())
	if err != nil {
		log.Println("Failed to get offers", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get offers"})
		return
	}
	c.JSON(http.StatusOK, offers)
}

// GetOfferByID handles GET requests and returns an offer by ID
func (h *OfferHandler) GetOfferByID(c *gin.Context) {
	id := c.Param("id")
	offer, err := h.offerService.GetOfferByID(c.Request.Context(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
			return
		}
		log.Println("Failed to get offer", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get offer"})
		return
	}
	c.JSON(http.StatusOK, offer)
}
