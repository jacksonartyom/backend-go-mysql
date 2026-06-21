package middleware

import (
	"backend-go-mysql/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my-secret-key")

func AuthMiddleware(jwtService *utils.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// ✅ ดึง userId จาก MapClaims
		userId := claims["userId"].(string)

		c.Set("userId", userId)

		c.Next()
	}
}
