package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// Middleware function for token verification
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			c.JSON(http.StatusForbidden, gin.H{"message": "No token provided!"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_JWT")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized!"})
			c.Abort()
			return
		}

		// If token is valid, you can access the claims
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// You can access the claims using claims["key"]
			// For example, req.UserId = claims["id"].(string)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized!"})
			c.Abort()
			return
		}
	}
}
