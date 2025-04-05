package moderation

import (
	"github.com/So-ham/Content-Moderation/internal/entities"
	"gorm.io/gorm"
)

type moderation struct {
	DB *gorm.DB
}

type Moderation interface {
	InsertModeration(moderation *entities.Moderation) (id uint, err error)
}

func New(db *gorm.DB) Moderation {
	return &moderation{DB: db}
}
