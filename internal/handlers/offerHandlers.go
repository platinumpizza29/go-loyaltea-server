package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"loyaltea-server/internal/models"
	"loyaltea-server/internal/services"
	"net/http"
	"os"
	"strings"

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

// MailchimpOfferRequest represents the expected POST payload for Mailchimp
// (Mailchimp expects JSON)
type MailchimpOfferRequest struct {
	SenderEmail string   `json:"sender_email" binding:"required,email"`
	Subject     string   `json:"subject"`
	Body        string   `json:"body"`
	Brand       string   `json:"brand"`
	Source      string   `json:"source"`
	Tags        []string `json:"tags"`
}

// subscribeToMailchimp subscribes an email to a Mailchimp list
func subscribeToMailchimp(email string) error {
	apiKey := os.Getenv("MAILCHIMP_API_KEY")
	listID := os.Getenv("MAILCHIMP_LIST_ID")
	if apiKey == "" || listID == "" {
		return fmt.Errorf("mailchimp API key or list ID not set in environment")
	}
	// Mailchimp API base URL (usX must match your API key's datacenter)
	datacenter := apiKey[strings.LastIndex(apiKey, "-")+1:]
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/members", datacenter, listID)

	payload := map[string]interface{}{
		"email_address": email,
		"status":        "subscribed",
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("anystring", apiKey) // Mailchimp uses anystring:apikey for basic auth

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("Mailchimp API error: %s", resp.Status)
	}
	return nil
}

// ReceiveOffer handles POST requests and subscribes sender to Mailchimp
func (h *OfferHandler) ReceiveOffer(c *gin.Context) {
	var req MailchimpOfferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Subscribe sender to Mailchimp
	if err := subscribeToMailchimp(req.SenderEmail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to subscribe to Mailchimp", "details": err.Error()})
		return
	}
	// Store offer in DB as before
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
	c.JSON(http.StatusOK, gin.H{"message": "Offer stored and user subscribed to Mailchimp successfully"})
}

// Add a simple GET endpoint to respond to Mailchimp webhook verification
func (h *OfferHandler) VerifyWebhook(c *gin.Context) {
	c.String(http.StatusOK, "Webhook endpoint verified")
}
