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

var externalCtxKey = &contextKey{"external"}

// Claims represents the JWT claims structure
type ExternalClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func ExternalCallbackMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization header. Must be a basic token.",
			})
			c.Abort()
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
		//getting the jwt key from env variables
		var jwtKey = []byte(os.Getenv("JWT_KEY"))
		// Verify the token
		parsedToken, err := jwt.ParseWithClaims(token, &ExternalClaims{}, func(token *jwt.Token) (interface{}, error) {
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
		externalclaims, ok := parsedToken.Claims.(*ExternalClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token claims",
			})
			c.Abort()
			return
		}

		// Store the user ID from the claims in the context
		ctx := context.WithValue(c.Request.Context(), externalCtxKey, externalclaims)
		c.Request = c.Request.WithContext(ctx)
		fmt.Println("authorized")
		c.Next()
	}
}

func ForExternalContext(ctx context.Context) *ExternalClaims {
	raw, _ := ctx.Value(externalCtxKey).(*ExternalClaims)
	return raw
}
