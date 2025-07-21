package models

type Stop struct {
	ID          string   `bson:"_id" json:"id"`            // MongoDB ID
	Name        string   `bson:"name" json:"name"`         // Store the store/stop name
	UserID      string   `bson:"user_id" json:"user_id"`   // User ID
	Address     string   `bson:"address" json:"address"`   // Optional human-readable address
	Location    GeoPoint `bson:"location" json:"location"` // GeoJSON point for geolocation
	IsCompleted bool     `bson:"isCompleted" json:"isCompleted"`
	Points      int      `bson:"points" json:"points"`
}

type GeoPoint struct {
	Type        string    `bson:"type" json:"type"`               // Always "Point"
	Coordinates []float64 `bson:"coordinates" json:"coordinates"` // [longitude, latitude]
}
