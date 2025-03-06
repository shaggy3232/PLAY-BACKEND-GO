package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define a struct for JWT claims
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Function to generate a JWT token
func GenerateJWT(userID string) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString([]byte(os.Getenv("jwtSecret")))
}
