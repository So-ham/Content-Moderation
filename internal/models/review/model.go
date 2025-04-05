package review

import (
	"gorm.io/gorm"
)

type review struct {
	DB *gorm.DB
}

type Review interface {
	AddReview(postID, userID uint, content string) (id uint, err error)
	UpdateFlagStatus(id uint, isFlagged bool) error
}

func New(db *gorm.DB) Review {
	return &review{DB: db}
}
