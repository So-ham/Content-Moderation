package services

import (
	"github.com/So-ham/Content-Moderation/internal/entities"
	"github.com/So-ham/Content-Moderation/internal/models"
)

// Service represents the service layer having
// all the services from all service packages
type service struct {
	model models.Model
}

// New creates a new instance of Service
func New(model *models.Model) Service {
	m := &service{model: *model}
	return m
}

type Service interface {
	Login(req *entities.UserLoginRequest) (*entities.UserResponse, string, error)

	Signup(req *entities.UserSignupRequest) (*entities.UserResponse, string, error)

	GetAllPosts() ([]entities.PostResponse, error)
}
