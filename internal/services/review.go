package services

import (
	"context"
	"fmt"
	"log"

	"github.com/So-ham/Content-Moderation/internal/entities"
	"github.com/So-ham/Content-Moderation/pkg/grpc/clients/notf"
	"github.com/So-ham/Content-Moderation/pkg/middlewares"
)

func (s *service) AddReview(ctx context.Context, req *entities.ReviewRequest) (resp *entities.ReviewResponse, err error) {

	curUser := middlewares.GetUserContext(ctx)

	commentID, err := s.model.Review.AddReview(req.PostID, curUser.UserID, req.Content)
	if err != nil {
		return nil, err
	}

	go func(commentID, userID, postID uint, content string) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered in review moderation goroutine: %v", r)
			}
		}()

		isFlagged, severity, err := s.tisane.CheckIfContentFlagged(content)
		if err != nil {
			log.Printf("Tisane moderation error: %v", err)
			return
		}

		if !isFlagged {
			return
		}

		go func(commentID uint) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in flag status update goroutine: %v", r)
				}
			}()
			if err := s.model.Review.UpdateFlagStatus(commentID, true); err != nil {
				log.Printf("Failed to update flag status for comment %d: %v", commentID, err)
			}
		}(commentID)

		go func(userID, postID uint, content, severity string) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in insert moderation goroutine: %v", r)
				}
			}()
			moderation := &entities.Moderation{
				UserID:   userID,
				Severity: severity,
				Review: entities.Review{
					UserID:    userID,
					PostID:    postID,
					Content:   content,
					Type:      "review",
					IsFlagged: true,
				},
			}
			if _, err := s.model.Moderation.InsertModeration(moderation); err != nil {
				log.Printf("Failed to insert moderation entry: %v", err)
			}
		}(userID, postID, content, severity)

		go func(userID uint, content, severity string) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in notification goroutine: %v", r)
				}
			}()
			fmt.Println("Sending review notification")
			_, err := s.notf.SendFlaggedNotification(context.Background(), &notf.SendFlaggedNotificationReq{
				UserID:   uint32(userID),
				Content:  content,
				Severity: severity,
			})
			if err != nil {
				log.Printf("Error sending review notification for userID %d: %v", userID, err)
			}
		}(userID, content, severity)

	}(commentID, curUser.UserID, req.PostID, req.Content)

	return &entities.ReviewResponse{
		ID:        commentID,
		Content:   req.Content,
		IsFlagged: false,
	}, nil
}
