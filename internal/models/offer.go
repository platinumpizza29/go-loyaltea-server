package models

import (
	"time"
)

type Offer struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	SenderEmail string    `bson:"sender_email" json:"sender_email"`         // Email of the user who forwarded it
	Subject     string    `bson:"subject" json:"subject"`                   // Subject line of the email
	Text        string    `bson:"text" json:"text"`                         // Plain text body
	Brand       string    `bson:"brand,omitempty" json:"brand,omitempty"`   // Optional: Parsed brand like "Zara", "Starbucks"
	Source      string    `bson:"source,omitempty" json:"source,omitempty"` // e.g., "email"
	Images      []string  `bson:"images" json:"images"`                     // Images from the email
	Ctas        []string  `bson:"ctas" json:"ctas"`                         // Call-to-action links
	Tags        []string  `bson:"tags,omitempty" json:"tags,omitempty"`     // Optional: e.g., ["discount", "clothing"]
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`               // When this offer was received
}
