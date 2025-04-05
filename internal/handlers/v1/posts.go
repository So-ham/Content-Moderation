package v1

import (
	"encoding/json"
	"net/http"
)

// HandlerV1 represents the version 1 handler

// GetAllPostsHandler handles the request to fetch all posts
func (h *handlerV1) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.GetAllPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
