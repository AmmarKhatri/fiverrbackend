package authhelpers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Claims represents the JWT claims structure
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// Check if the Authorization header starts with "Basic"
		if !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization header. Must be a basic token.",
			})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Basic ")
		// Allow unauthenticated users in to access public endpoints or get rejected by private ones
		if token == "" {
			//allowing no token but with Basic header into the API
			c.Next()
			return
		}
		//getting the jwt key from env variables
		var jwtKey = []byte(os.Getenv("JWT_KEY"))
		// Verify the token
		parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		// Extract the claims from the token
		claims, ok := parsedToken.Claims.(*Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		// Store the user ID from the claims in the context
		ctx := context.WithValue(c.Request.Context(), userCtxKey, claims)
		c.Request = c.Request.WithContext(ctx)
		fmt.Println("authorized")
		c.Next()
	}
}
func ForContext(ctx context.Context) *Claims {
	raw, _ := ctx.Value(userCtxKey).(*Claims)
	return raw
}
