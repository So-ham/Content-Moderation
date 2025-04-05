package models

import (
	"github.com/So-ham/Content-Moderation/internal/models/post"
	"github.com/So-ham/Content-Moderation/internal/models/user"

	"gorm.io/gorm"
)

type Model struct {
	User user.User
	Post post.Post
}

// New creates a new instance of Model
func New(gdb *gorm.DB) *Model {
	return &Model{
		User: user.New(gdb),
		Post: post.New(gdb),
	}
}
