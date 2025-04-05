package post

import (
	"github.com/So-ham/Content-Moderation/internal/entities"

	"gorm.io/gorm"
)

type post struct {
	DB *gorm.DB
}

type Post interface {
	GetAllPosts() ([]entities.Post, error)
}

func New(db *gorm.DB) Post {
	return &post{DB: db}
}
