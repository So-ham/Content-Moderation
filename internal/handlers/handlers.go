package handlers

import (
	v1 "github.com/So-ham/Content-Moderation/internal/handlers/v1"
	"github.com/So-ham/Content-Moderation/internal/services"
)

type Handler struct {
	V1 v1.HandlerV1
}

// New creates a new instance of handlers
func New(service services.Service) *Handler {
	return &Handler{
		V1: v1.New(service),
	}

}
