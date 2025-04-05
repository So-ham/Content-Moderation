package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/So-ham/Content-Moderation/internal/db/postgres"
	"github.com/So-ham/Content-Moderation/internal/handlers"
	"github.com/So-ham/Content-Moderation/internal/models"
	"github.com/So-ham/Content-Moderation/internal/services"
	"github.com/So-ham/Content-Moderation/internal/web/rest"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	db := postgres.Connect()

	model := models.New(db)
	fmt.Println("Model layer initialized")

	service := services.New(model)
	fmt.Println("Service layer initialized")

	handler := handlers.New(service)
	fmt.Println("Handler layer initialized")

	r := rest.NewRouter(handler)
	fmt.Println("Routers loaded")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	corsHandler := c.Handler(r)

	// Start server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", corsHandler)
}
