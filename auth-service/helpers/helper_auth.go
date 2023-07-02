package helpers

import (
	"auth-service/models"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken() (string, error) {
	// Create the claims
	claims := &models.Claims{
		ID: 1,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
