package services

import (
	"context"

	"github.com/So-ham/Content-Moderation/internal/entities"
	"github.com/So-ham/Content-Moderation/internal/models"
	"github.com/So-ham/Content-Moderation/pkg/grpc/clients/notf"
	"github.com/So-ham/Content-Moderation/pkg/tisane"
)

// Service represents the service layer having
// all the services from all service packages
type service struct {
	model  models.Model
	tisane tisane.TiSane
	notf   notf.NotfServiceClient
}

// New creates a new instance of Service
func New(model *models.Model) Service {
	m := &service{model: *model, tisane: tisane.New(), notf: notf.NewClient()}
	return m
}

type Service interface {
	Login(req *entities.UserLoginRequest) (*entities.UserResponse, string, error)

	Signup(ctx context.Context, req *entities.UserSignupRequest) (*entities.UserResponse, string, error)

	GetAllPosts() ([]entities.PostResponse, error)

	AddComment(ctx context.Context, req *entities.CommentRequest) (*entities.CommentResponse, error)

	AddReview(ctx context.Context, req *entities.ReviewRequest) (resp *entities.ReviewResponse, err error)
}
