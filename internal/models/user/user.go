package user

import (
	"errors"

	"github.com/So-ham/Content-Moderation/internal/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Create creates a new user in the database
func (m *user) Create(user *entities.User) error {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

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
