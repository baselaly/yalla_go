package middleware

import (
	"net/http"
	"yalla_go/jwt"

	"github.com/gin-gonic/gin"
)

// Auth check user authentication
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := jwt.VerifyToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorzied"})
			c.Abort()
		}

		c.Next()
	}
}
