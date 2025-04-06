package user

import (
	"context"
	"errors"

	"github.com/So-ham/Content-Moderation/internal/entities"

	"gorm.io/gorm"
)

// Create creates a new user in the database
func (m *user) Create(ctx context.Context, user *entities.User) error {

	result := m.DB.Create(user)
	return result.Error
}

// FindByEmail finds a user by email
func (m *user) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := m.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
