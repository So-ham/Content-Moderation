package moderation

import "github.com/So-ham/Content-Moderation/internal/entities"

// InsertModeration adds a new moderation record to the database using GORM
func (m *moderation) InsertModeration(moderation *entities.Moderation) (id uint, err error) {

	err = m.DB.Create(moderation).Error
	return moderation.ID, err
}
