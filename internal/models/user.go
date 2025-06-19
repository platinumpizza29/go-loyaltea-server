package models

import (
	"time"
)

// User represents the user model
type User struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"-"`
	Name      string    `bson:"name" json:"name"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
