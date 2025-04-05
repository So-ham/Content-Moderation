package services

import "github.com/So-ham/Content-Moderation/internal/entities"

func (s *service) GetAllPosts() ([]entities.PostResponse, error) {
	posts, err := s.model.Post.GetAllPosts()
	if err != nil {
		return nil, err
	}

	var postResponses []entities.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, entities.PostResponse{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			User:      post.User,
			Reviews:   post.Reviews,
		})
	}
	//

	return postResponses, nil
}
