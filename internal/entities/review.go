package entities

import (
	"time"
)

// Review represents a user review or comment with moderation status
type Review struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	PostID      uint       `json:"post_id" gorm:"not null;index"`
	UserID      uint       `json:"user_id" gorm:"not null;index"`
	Content     string     `json:"content" gorm:"type:text"`
	IsFlagged   bool       `json:"is_flagged" gorm:"default:false;index"`
	ModeratedAt *time.Time `json:"moderated_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at" gorm:"index"`
	UpdatedAt   time.Time  `json:"updated_at"`
	User        User       `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// ReviewRequest represents the request body for creating a review or comment
type ReviewRequest struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

// ReviewResponse represents the response body for review-related operations
type ReviewResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsFlagged bool      `json:"is_flagged"`
	Severity  string    `json:"severity,omitempty"`
	User      User      `json:"user"`
	Post      Post      `json:"post"`
}
