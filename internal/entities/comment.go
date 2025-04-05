package entities

import "time"

// Comment represents a comment made by a user on a post
// Updated to include moderation fields similar to Review
type Comment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PostID      uint      `json:"post_id" gorm:"not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	Content     string    `json:"content" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsFlagged   bool      `json:"is_flagged" gorm:"default:false;index"`
	ModeratedAt time.Time `json:"moderated_at,omitempty"`
	User        User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Post        Post      `json:"post" gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE"`
}

// CommentRequest represents the request body for creating a comment
// Updated to include moderation fields
type CommentRequest struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

// CommentResponse represents the response body for comment-related operations
// Updated to include moderation fields
type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsFlagged bool      `json:"is_flagged"`
	Severity  string    `json:"severity,omitempty"`
	User      User      `json:"user"`
	Post      Post      `json:"post"`
}
