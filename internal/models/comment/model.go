package comment

import (
	"gorm.io/gorm"
)

type comment struct {
	DB *gorm.DB
}

type Comment interface {
	AddComment(postID, userID uint, content string) (id uint, err error)
	UpdateFlagStatus(id uint, isFlagged bool) error
}

func New(db *gorm.DB) Comment {
	return &comment{DB: db}
}
