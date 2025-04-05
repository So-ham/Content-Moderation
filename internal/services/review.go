package services

import (
	"context"

	"github.com/So-ham/Content-Moderation/internal/entities"
	"github.com/So-ham/Content-Moderation/pkg/middlewares"
)

func (s *service) AddReview(ctx context.Context, req *entities.ReviewRequest) (resp *entities.ReviewResponse, err error) {

	curUser := middlewares.GetUserContext(ctx)

	commentID, err := s.model.Review.AddReview(req.PostID, curUser.UserID, req.Content)
	if err != nil {
		return nil, err
	}
	// Check content moderation using Tisane
	isFlagged, severity, err := s.tisane.CheckIfContentFlagged(req.Content)
	if err != nil {
		return nil, err
	}

	if isFlagged {
		// Update comment's flagged status
		err := s.model.Review.UpdateFlagStatus(commentID, true)
		if err != nil {
			return nil, err
		}

		// Create moderation entry
		moderation := &entities.Moderation{
			UserID:   curUser.UserID,
			Severity: severity,
			Review: entities.Review{
				UserID:    curUser.UserID,
				PostID:    req.PostID,
				Content:   req.Content,
				Type:      "review",
				IsFlagged: true,
			},
		}

		_, err = s.model.Moderation.InsertModeration(moderation)
		if err != nil {
			return nil, err
		}

		return &entities.ReviewResponse{
			ID:        commentID,
			Content:   req.Content,
			IsFlagged: true,
			Severity:  severity,
		}, nil

	}

	return &entities.ReviewResponse{
		ID:        commentID,
		Content:   req.Content,
		IsFlagged: false,
	}, nil
}
