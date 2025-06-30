package models

import (
	"time"
)

type Offer struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	SenderEmail string    `bson:"senderEmail" json:"senderEmail"`           // Email of the user who forwarded it
	Subject     string    `bson:"subject" json:"subject"`                   // Subject line of the email
	Body        string    `bson:"body" json:"body"`                         // Plain text body
	Brand       string    `bson:"brand,omitempty" json:"brand,omitempty"`   // Optional: Parsed brand like "Zara", "Starbucks"
	Source      string    `bson:"source,omitempty" json:"source,omitempty"` // e.g., "email"
	Images      []string  `bson:"images" json:"images"`                     // e.g., "email"
	Links       []string  `bson:"links" json:"links"`                       // e.g., "email"
	Tags        []string  `bson:"tags,omitempty" json:"tags,omitempty"`     // Optional: e.g., ["discount", "clothing"]
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`               // When this offer was received
}
