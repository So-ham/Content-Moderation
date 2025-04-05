package v1

import (
	"encoding/json"
	"net/http"
)

// GetAllPostsHandler retrieves all posts from the system and returns them as JSON.
func (h *handlerV1) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.GetAllPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
