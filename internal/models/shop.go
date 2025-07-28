package models

import "time"

type GeoPoint struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

type Shop struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Brand       string    `bson:"brand" json:"brand"`
	Address     string    `bson:"address" json:"address"`
	Location    GeoPoint  `bson:"location" json:"location"`
	PointsValue int       `bson:"points_value" json:"points_value"` // Points awarded for visiting
	Category    string    `bson:"category" json:"category"`         // e.g., "coffee", "clothing", "grocery"
	IsActive    bool      `bson:"is_active" json:"is_active"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
