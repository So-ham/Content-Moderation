package user

import (
	"context"

	"github.com/So-ham/Content-Moderation/internal/entities"

	"gorm.io/gorm"
)

type user struct {
	DB *gorm.DB
}

type User interface {
	Create(ctx context.Context, user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}

func New(db *gorm.DB) User {
	return &user{DB: db}
}
