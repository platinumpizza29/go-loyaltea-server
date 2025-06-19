package db

import (
	"context"
	"fmt"
	"loyaltea-server/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

// UserModel handles database operations for users
type UserModel struct {
	collection *mongo.Collection
}

// NewUserModel creates a new UserModel instance
func NewUserModel(db *mongo.Database) *UserModel {
	return &UserModel{
		collection: db.Collection("users"),
	}
}

// Create creates a new user
func (m *UserModel) Create(ctx context.Context, user *models.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Password = string(hashedPassword)

	// Insert the user
	_, err = m.collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}

	return nil
}

// FindByEmail finds a user by email
func (m *UserModel) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := m.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByID finds a user by ID
func (m *UserModel) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update updates a user
func (m *UserModel) Update(ctx context.Context, user *models.User) error {
	user.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": user.UpdatedAt,
		},
	}

	_, err := m.collection.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		update,
	)
	return err
}

// Delete deletes a user
func (m *UserModel) Delete(ctx context.Context, id string) error {
	_, err := m.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// VerifyPassword checks if the provided password matches the user's password
func (m *UserModel) VerifyPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
