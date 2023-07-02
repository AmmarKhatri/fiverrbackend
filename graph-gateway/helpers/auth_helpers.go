package helpers

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
		token := c.Request.Header.Get("Authorization")
		fmt.Println(token, "is my token")

		// Allow unauthenticated users in to access public endpoints or get rejected by private ones
		if token == "" {
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
		ctx := context.WithValue(c.Request.Context(), userCtxKey, claims.UserID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
		fmt.Println("authorized")
		c.Next()
	}
}
func ForContext(ctx context.Context) *Claims {
	raw, _ := ctx.Value(userCtxKey).(*Claims)
	return raw
}
