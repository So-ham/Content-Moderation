package v1

import (
	"encoding/json"
	"net/http"

	"github.com/So-ham/Content-Moderation/internal/entities"
)

// AddReviewHandler processes requests to add content moderation reviews.
func (h *handlerV1) AddReviewHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	if r.Body == http.NoBody {
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	var req entities.ReviewRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	commentResponse, err := h.Service.AddReview(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Assume we have a function to write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentResponse)
}
