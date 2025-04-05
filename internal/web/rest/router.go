package rest

import (
	"net/http"

	"github.com/So-ham/Content-Moderation/internal/handlers"

	"github.com/gorilla/mux"
)

// NewRouter returns a new router instance with configured routes
func NewRouter(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	// Auth endpoints
	router.HandleFunc("/auth/signup", h.V1.SignupHandler).Methods("POST")
	router.HandleFunc("/auth/login", h.V1.LoginHandler).Methods("POST")

	// Post endpoints
	router.Handle("/posts", JWTMiddleware(http.HandlerFunc(h.V1.GetAllPostsHandler))).Methods("GET")

	return router
}
