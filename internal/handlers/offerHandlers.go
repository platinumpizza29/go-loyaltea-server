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

// GetOffers godoc
// @Summary      Get all offers
// @Description  Retrieves a list of all available offers
// @Tags         offers
// @Produce      json
// @Success      200 {array} models.Offer
// @Failure      500 {object} map[string]string
// @Router       /offers [get]
func (h *OfferHandler) GetOffers(c *gin.Context) {
	offers, err := h.offerService.GetOffers(c.Request.Context())
	if err != nil {
		log.Println("Failed to get offers", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get offers"})
		return
	}
	c.JSON(http.StatusOK, offers)
}

// GetOfferByID godoc
// @Summary      Get offer by ID
// @Description  Retrieves a single offer by its ID
// @Tags         offers
// @Produce      json
// @Param        id   path      string  true  "Offer ID"
// @Success      200  {object}  models.Offer
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /offers/{id} [get]
func (h *OfferHandler) GetOfferByID(c *gin.Context) {
	id := c.Param("id")
	offer, err := h.offerService.GetOfferByID(c.Request.Context(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, map[string]string{"error": "Offer not found"})
			return
		}
		log.Println("Failed to get offer", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get offer"})
		return
	}
	c.JSON(http.StatusOK, offer)
}
