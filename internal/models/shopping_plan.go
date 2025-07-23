package models

import "time"

type ShoppingPlan struct {
	ID          string     `bson:"_id,omitempty" json:"id"`
	UserID      string     `bson:"user_id" json:"user_id"`
	Name        string     `bson:"name" json:"name"` // e.g., "Weekend Shopping"
	Description string     `bson:"description" json:"description"`
	Stops       []PlanStop `bson:"stops" json:"stops"`
	IsCompleted bool       `bson:"is_completed" json:"is_completed"`
	CreatedAt   time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at" json:"updated_at"`
}

type PlanStop struct {
	ShopID       string     `bson:"shop_id" json:"shop_id"`
	IsVisited    bool       `bson:"is_visited" json:"is_visited"`
	VisitedAt    *time.Time `bson:"visited_at,omitempty" json:"visited_at,omitempty"`
	PointsEarned int        `bson:"points_earned" json:"points_earned"`
}
