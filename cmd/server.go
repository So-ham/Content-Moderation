package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/So-ham/Content-Moderation/internal/db/postgres"
	"github.com/So-ham/Content-Moderation/internal/handlers"
	"github.com/So-ham/Content-Moderation/internal/models"
	"github.com/So-ham/Content-Moderation/internal/services"
	"github.com/So-ham/Content-Moderation/internal/web/rest"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	if os.Getenv("ENV") == "" { // load Env only on production
		if err := godotenv.Load(); err != nil {
			log.Fatalln("Error loading env file", err)
		}
	}

	v := validator.New()

	db := postgres.Connect()

	model := models.New(db)
	fmt.Println("Model layer initialized")

	service := services.New(model)
	fmt.Println("Service layer initialized")

	handler := handlers.New(service, v)
	fmt.Println("Handler layer initialized")

	r := rest.NewRouter(handler)
	fmt.Println("Routers loaded")

	allowedOrigins := []string{"http://localhost:3000"}
	if os.Getenv("ENV") == "prod" {
		allowedOrigins = []string{"https://content-moderation-frontend.vercel.app"}
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	corsHandler := c.Handler(r)

	// Start server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", corsHandler)
}
