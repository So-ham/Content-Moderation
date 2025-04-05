package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/So-ham/Content-Moderation/internal/entities"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		log.Println("Authorization:", r.Header.Get("Authorization"))

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		claims, err := ValidateToken(tokenParts[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetUserContext(ctx context.Context) *entities.CustomClaims {
	// Retrieve the claims from the context
	claims, _ := ctx.Value("claims").(*entities.CustomClaims)

	return claims
}
