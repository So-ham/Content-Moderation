package services

import (
	"errors"
	"time"

	"github.com/So-ham/Content-Moderation/internal/auth"
	"github.com/So-ham/Content-Moderation/internal/entities"

	"golang.org/x/crypto/bcrypt"
)

// Signup handles user registration
func (s *service) Signup(req *entities.UserSignupRequest) (*entities.UserResponse, string, error) {
	// Check if user already exists
	existingUser, err := s.model.User.FindByEmail(req.Email)
	if err != nil {
		return nil, "", err
	}
	if existingUser != nil {
		return nil, "", errors.New("user with this email already exists")
	}

	// Create new user
	newUser := &entities.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save user to database
	if err = s.model.User.Create(newUser); err != nil {
		return nil, "", err
	}

	// Generate JWT token
	token, err := auth.GenerateToken(newUser)
	if err != nil {
		return nil, "", err
	}

	// Prepare response
	response := &entities.UserResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}

	return response, token, nil
}

// Login handles user authentication
func (s *service) Login(req *entities.UserLoginRequest) (*entities.UserResponse, string, error) {
	// Find user by email
	user, err := s.model.User.FindByEmail(req.Email)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	// Prepare response
	response := &entities.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return response, token, nil
}
