package handlers

import (
	"loyaltea-server/internal/models"
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

// MailgunOfferRequest represents the expected POST payload from Mailgun
// (Mailgun sends form data, not JSON)
type MailgunOfferRequest struct {
	SenderEmail string   `form:"sender"`
	Subject     string   `form:"subject"`
	Body        string   `form:"body-plain"`
	Brand       string   `form:"brand"`
	Source      string   `form:"source"`
	Tags        []string `form:"tags[]"`
}

// ReceiveOffer handles POST requests from Mailgun webhooks
func (h *OfferHandler) ReceiveOffer(c *gin.Context) {
	var req MailgunOfferRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	offer := &models.Offer{
		SenderEmail: req.SenderEmail,
		Subject:     req.Subject,
		Body:        req.Body,
		Brand:       req.Brand,
		Source:      req.Source,
		Tags:        req.Tags,
	}
	if err := h.offerService.CreateOffer(c.Request.Context(), offer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store offer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Offer stored successfully"})
}
