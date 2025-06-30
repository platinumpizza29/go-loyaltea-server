package handlers

import (
	"log"
	"loyaltea-server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OfferHandler struct {
	offerService *services.OfferService
}

func NewOfferHandler(offerService *services.OfferService) *OfferHandler {
	return &OfferHandler{
		offerService: offerService,
	}
}

// MailgunOfferRequest represents the expected POST payload for Mailgun
// (expects JSON)
type MailgunOfferRequest struct {
	SenderEmail string   `json:"sender_email" binding:"required,email"`
	Subject     string   `json:"subject"`
	Body        string   `json:"body"`
	Brand       string   `json:"brand"`
	Source      string   `json:"source"`
	Tags        []string `json:"tags"`
}

// Add a simple GET endpoint to respond to Mailchimp webhook verification
func (h *OfferHandler) VerifyWebhook(c *gin.Context) {
	c.String(http.StatusOK, "Webhook endpoint verified")
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get offer"})
		return
	}
	c.JSON(http.StatusOK, offer)
}
