package v1

import (
	"net/http"

	"github.com/So-ham/Content-Moderation/internal/services"
	"github.com/go-playground/validator/v10"
)

// handlerV1 represents the v1 handler
type handlerV1 struct {
	Service  services.Service
	Validate *validator.Validate
}

// HandlerV1 has handlers list
type HandlerV1 interface {
	SignupHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)

	GetAllPostsHandler(w http.ResponseWriter, r *http.Request)

	AddCommentHandler(w http.ResponseWriter, r *http.Request)

	AddReviewHandler(w http.ResponseWriter, r *http.Request)
}

// New creates an instance of handlers for V1
func New(s services.Service, v *validator.Validate) HandlerV1 {
	return &handlerV1{Service: s, Validate: v}
}
