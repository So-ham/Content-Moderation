package entities

import (
	"time"
)

// Review represents a user review or comment with moderation status
type Review struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"not null;index"`
	PostID          uint      `json:"post_id" gorm:"not null;index"`
	Content         string    `json:"content" gorm:"not null;type:text"`
	Type            string    `json:"type" gorm:"not null;type:varchar(20);check:type IN ('review', 'comment')"` // 'review' or 'comment'
	IsFlagged       bool      `json:"is_flagged" gorm:"default:false;index"`
	ModerationScore float64   `json:"moderation_score,omitempty" gorm:"type:decimal(3,2);default:0"`
	ModeratedAt     time.Time `json:"moderated_at,omitempty"`
	CreatedAt       time.Time `json:"created_at" gorm:"index"`
	UpdatedAt       time.Time `json:"updated_at"`
	User            User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Post            Post      `json:"post" gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE"`
}

// ReviewRequest represents the request body for creating a review or comment
type ReviewRequest struct {
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
	Type    string `json:"type" binding:"required,oneof=review comment"`
}

// ReviewResponse represents the response body for review-related operations
type ReviewResponse struct {
	ID              uint      `json:"id"`
	Content         string    `json:"content"`
	Type            string    `json:"type"`
	IsFlagged       bool      `json:"is_flagged"`
	ModerationScore float64   `json:"moderation_score,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	User            User      `json:"user"`
	Post            Post      `json:"post"`
}
