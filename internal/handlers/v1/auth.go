package v1

import (
	"encoding/json"
	"net/http"

	"github.com/So-ham/Content-Moderation/internal/entities"
)

// SignupHandler handles user registration
func (h *handlerV1) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req entities.UserSignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.Validate.Struct(req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call service
	user, token, err := h.Service.Signup(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare response
	response := struct {
		User  *entities.UserResponse `json:"user"`
		Token string                 `json:"token"`
	}{
		User:  user,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LoginHandler handles user authentication
func (h *handlerV1) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	var req entities.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call service
	user, token, err := h.Service.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Prepare response
	response := struct {
		User  *entities.UserResponse `json:"user"`
		Token string                 `json:"token"`
	}{
		User:  user,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
