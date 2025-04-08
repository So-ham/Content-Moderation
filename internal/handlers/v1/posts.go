package v1

import (
	"encoding/json"
	"net/http"
)

// GetAllPostsHandler retrieves all posts from the system and returns them as JSON.
func (h *handlerV1) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	posts, err := h.Service.GetAllPosts(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
