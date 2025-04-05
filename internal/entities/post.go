package entities

import (
	"time"
)

// Post represents a social media post created by users
type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"` // Index for faster queries
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	Reviews   []Review  `json:"reviews,omitempty" gorm:"foreignKey:PostID"`
}

// PostRequest represents the request body for creating a post
type PostRequest struct {
	Content string `json:"content" validate:"required"`
}

// PostResponse represents the response body for post-related operations
type PostResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
	Reviews   []Review  `json:"reviews,omitempty"`
}
