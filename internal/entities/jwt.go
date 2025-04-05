package entities

import "github.com/golang-jwt/jwt/v4"

// CustomClaims represents the claims in the JWT token
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
