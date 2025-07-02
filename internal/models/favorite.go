package models

import (
	"time"
)

type Favorite struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	UserID       string    `json:"user_id" bson:"user_id"`
	OfferId      string    `json:"offer_id" bson:"offer_id"`
	OfferCreated time.Time `json:"offer_created" bson:"offer_created"`
}
