package models

import (
	"github.com/So-ham/Content-Moderation/internal/models/comment"
	"github.com/So-ham/Content-Moderation/internal/models/moderation"
	"github.com/So-ham/Content-Moderation/internal/models/post"
	"github.com/So-ham/Content-Moderation/internal/models/review"
	"github.com/So-ham/Content-Moderation/internal/models/user"

	"gorm.io/gorm"
)

type Model struct {
	User       user.User
	Post       post.Post
	Comment    comment.Comment
	Moderation moderation.Moderation
	Review     review.Review
}

// New creates a new instance of Model
func New(gdb *gorm.DB) *Model {
	return &Model{
		User:       user.New(gdb),
		Post:       post.New(gdb),
		Comment:    comment.New(gdb),
		Moderation: moderation.New(gdb),
		Review:     review.New(gdb),
	}
}
