package entities

import (
	"time"
)

// Moderation represents a flagged content entry for moderation
type Moderation struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ReviewID  uint      `json:"review_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null;index"` // Index for faster queries
	Severity  string    `json:"reason" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Review    Review    `json:"review" gorm:"foreignKey:ReviewID"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}

// ModerationResponse represents the response for moderation-related operations
type ModerationResponse struct {
	ID         uint      `json:"id"`
	Reason     string    `json:"reason"`
	Status     string    `json:"status"`
	Notified   bool      `json:"notified"`
	NotifiedAt time.Time `json:"notified_at,omitempty"`
	Review     Review    `json:"review"`
}
