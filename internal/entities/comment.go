package entities

import "time"

// Comment represents a comment made by a user on a post
// Updated to include moderation fields similar to Review
type Comment struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	PostID      uint       `json:"post_id" gorm:"not null;index"`
	UserID      uint       `json:"user_id" gorm:"not null;index"`
	Content     string     `json:"content" gorm:"not null;type:text"`
	IsFlagged   bool       `json:"is_flagged" gorm:"default:false;index"`
	ModeratedAt *time.Time `json:"moderated_at,omitempty"` // Pointer allows NULL
	CreatedAt   time.Time  `json:"created_at" gorm:"index"`
	UpdatedAt   time.Time  `json:"updated_at"`
	User        User       `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
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
