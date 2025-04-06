package comment

import (
	"fmt"
	"time"

	"github.com/So-ham/Content-Moderation/internal/entities"
)

// AddComment adds a comment to a post in the database
func (m *comment) AddComment(postID, userID uint, content string) (id uint, err error) {
	comment := &entities.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}
	err = m.DB.Model(&entities.Comment{}).Create(comment).Error
	return comment.ID, err
}

// UpdateFlagStatus updates the isFlagged status of a comment
func (m *comment) UpdateFlagStatus(id uint, isFlagged bool) error {
	err := m.DB.Model(&entities.Comment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_flagged":   isFlagged,
		"moderated_at": time.Now(),
	})
	if err.Error != nil {
		return fmt.Errorf("failed to update comment: %v", err.Error)
	}
	return nil
}
