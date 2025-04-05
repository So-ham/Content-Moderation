package post

import (
	"github.com/So-ham/Content-Moderation/internal/entities"
)

// GetAllPosts retrieves all posts from the database
func (m *post) GetAllPosts() ([]entities.Post, error) {
	var posts []entities.Post
	result := m.DB.Find(&posts)
	return posts, result.Error
}
