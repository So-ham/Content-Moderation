package v1

import (
	"encoding/json"
	"net/http"

	"github.com/So-ham/Content-Moderation/internal/entities"
)

// AddCommentHandler processes requests to add new comments to posts.
func (h *handlerV1) AddCommentHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	if r.Body == http.NoBody {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	var req entities.CommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	commentResponse, err := h.Service.AddComment(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Assume we have a function to write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentResponse)
}
