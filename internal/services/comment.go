package services

import (
	"context"
	"fmt"

	"github.com/So-ham/Content-Moderation/internal/entities"
	"github.com/So-ham/Content-Moderation/pkg/grpc/clients/notf"
	"github.com/So-ham/Content-Moderation/pkg/middlewares"
)

func (s *service) AddComment(ctx context.Context, req *entities.CommentRequest) (resp *entities.CommentResponse, err error) {
	curUser := middlewares.GetUserContext(ctx)

	commentID, err := s.model.Comment.AddComment(req.PostID, curUser.UserID, req.Content)
	if err != nil {
		return nil, err
	}

	// Run moderation asynchronously
	go func(commentID uint, content string, userID uint, postID uint) {
		isFlagged, severity, err := s.tisane.CheckIfContentFlagged(content)
		if err != nil {
			// You might want to log this error
			return
		}

		if isFlagged {
			// Update flagged status
			go func(commentID uint) {
				err := s.model.Comment.UpdateFlagStatus(commentID, true)
				if err != nil {
					fmt.Printf("Error updating comment flag status for commentID %d: %v\n", commentID, err)
				}
			}(commentID)

			go func(userID uint, postID uint, content string, severity string) {
				moderation := &entities.Moderation{
					UserID:   userID,
					Severity: severity,
					Content:  content,
				}
				_, err := s.model.Moderation.InsertModeration(moderation)
				if err != nil {
					fmt.Printf("Error inserting moderation entry for userID %d: %v\n", userID, err)
				}
			}(curUser.UserID, req.PostID, req.Content, severity)

			go func(userID uint, content, severity string) {
				fmt.Println("Sending comment notification")
				_, err := s.notf.SendFlaggedNotification(context.Background(), &notf.SendFlaggedNotificationReq{
					UserID:   uint32(userID),
					Content:  content,
					Severity: severity,
				})
				if err != nil {
					fmt.Printf("Error sending comment notification for userID %d: %v\n", userID, err)
				}
			}(curUser.UserID, req.Content, severity)

		}
	}(commentID, req.Content, curUser.UserID, req.PostID)

	// Return immediately (assume not flagged yet)
	return &entities.CommentResponse{
		ID:        commentID,
		Content:   req.Content,
		IsFlagged: false,
	}, nil
}
