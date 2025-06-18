package models

import (
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Offer struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderEmail string             `bson:"senderEmail" json:"senderEmail"`           // Email of the user who forwarded it
	Subject     string             `bson:"subject" json:"subject"`                   // Subject line of the email
	Body        string             `bson:"body" json:"body"`                         // Plain text body
	Brand       string             `bson:"brand,omitempty" json:"brand,omitempty"`   // Optional: Parsed brand like "Zara", "Starbucks"
	Source      string             `bson:"source,omitempty" json:"source,omitempty"` // e.g., "email"
	Tags        []string           `bson:"tags,omitempty" json:"tags,omitempty"`     // Optional: e.g., ["discount", "clothing"]
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`               // When this offer was received
}

// OfferModel handles database operations for offers
// Similar to UserModel for users
type OfferModel struct {
	collection *mongo.Collection
}

// NewOfferModel creates a new OfferModel instance
func NewOfferModel(db *mongo.Database) *OfferModel {
	return &OfferModel{
		collection: db.Collection("offers"),
	}
}

// Create inserts a new offer into the collection
func (m *OfferModel) Create(ctx context.Context, offer *Offer) error {
	offer.CreatedAt = time.Now()
	result, err := m.collection.InsertOne(ctx, offer)
	if err != nil {
		return err
	}
	offer.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
