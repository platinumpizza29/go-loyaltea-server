package services

import (
	"context"
	"errors"
	"regexp"
	"time"

	"loyaltea-server/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidPassword    = errors.New("password must be at least 8 characters long")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// UserService handles business logic for user operations
type UserService struct {
	userModel *models.UserModel
}

// NewUserService creates a new UserService instance
func NewUserService(userModel *models.UserModel) *UserService {
	return &UserService{
		userModel: userModel,
	}
}

// RegisterUser handles user registration
func (s *UserService) RegisterUser(email, password, name string) (*models.User, error) {
	// Validate email
	if !isValidEmail(email) {
		return nil, ErrInvalidEmail
	}

	// Validate password
	if len(password) < 8 {
		return nil, ErrInvalidPassword
	}

	// Check if user already exists
	existingUser, err := s.userModel.FindByEmail(context.TODO(), email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrEmailExists
	}

	// Create new user
	user := &models.User{
		Email:    email,
		Password: password,
		Name:     name,
	}

	err = s.userModel.Create(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// LoginUser handles user login
func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	// Find user by email
	user, err := s.userModel.FindByEmail(context.TODO(), email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	// Verify password
	if !s.userModel.VerifyPassword(user, password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	user, err := s.userModel.FindByID(context.TODO(), objectID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

// UpdateUser updates user information
func (s *UserService) UpdateUser(id string, email, name string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Get existing user
	user, err := s.userModel.FindByID(context.TODO(), objectID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	// Validate email if it's being changed
	if email != user.Email {
		if !isValidEmail(email) {
			return nil, ErrInvalidEmail
		}

		// Check if new email already exists
		existingUser, err := s.userModel.FindByEmail(context.TODO(), email)
		if err != nil {
			return nil, err
		}
		if existingUser != nil {
			return nil, ErrEmailExists
		}
	}

	// Update user fields
	user.Email = email
	user.Name = name
	user.UpdatedAt = time.Now()

	err = s.userModel.Update(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Check if user exists
	user, err := s.userModel.FindByID(context.TODO(), objectID)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	return s.userModel.Delete(context.TODO(), objectID)
}

// isValidEmail validates email format
func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
