package rest

import (
	"net/http"

	"github.com/So-ham/Content-Moderation/internal/handlers"
	"github.com/So-ham/Content-Moderation/pkg/middlewares"

	"github.com/gorilla/mux"
)

// NewRouter returns a new router instance with configured routes
func NewRouter(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	// Auth endpoints
	router.HandleFunc("/auth/signup", h.V1.SignupHandler).Methods("POST")
	router.HandleFunc("/auth/login", h.V1.LoginHandler).Methods("POST")

	// Post endpoints
	router.Handle("/posts", middlewares.JWTMiddleware(http.HandlerFunc(h.V1.GetAllPostsHandler))).Methods("GET")
	router.Handle("/comment", middlewares.JWTMiddleware(http.HandlerFunc(h.V1.AddCommentHandler))).Methods("POST")

	router.Handle("/review", middlewares.JWTMiddleware(http.HandlerFunc(h.V1.AddReviewHandler))).Methods("POST")

	// router.HandleFunc("/comment", h.V1.AddCommentHandler).Methods("POST")

	return router
}
