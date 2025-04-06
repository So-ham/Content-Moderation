package review

import (
	"fmt"
	"time"

	"github.com/So-ham/Content-Moderation/internal/entities"
)

// AddComment adds a review to a post in the database
func (m *review) AddReview(postID, userID uint, content string) (id uint, err error) {
	review := &entities.Review{
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}
	err = m.DB.Model(&entities.Review{}).Create(review).Error
	return review.ID, err
}

// UpdateFlagStatus updates the isFlagged status of a review
func (m *review) UpdateFlagStatus(id uint, isFlagged bool) error {
	err := m.DB.Model(&entities.Review{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_flagged":   isFlagged,
		"moderated_at": time.Now(),
	})
	if err.Error != nil {
		return fmt.Errorf("failed to update review: %v", err.Error)
	}
	return nil
}
