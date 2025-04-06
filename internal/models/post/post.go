package post

import (
	"github.com/So-ham/Content-Moderation/internal/entities"
)

// GetAllPosts retrieves all posts from the database
func (m *post) GetAllPosts() ([]entities.Post, error) {
	var posts []entities.Post

	// Use Preload to load all relationships
	result := m.DB.
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Reviews").
		Preload("Reviews.User"). // Load all Reviews // Load User for each Review
		Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}
