package middleware

import (
	"hey-notes-api/helpers"
	"hey-notes-api/internal/api/http/exception"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len("Bearer "):]
		if tokenString == "" {
			exception.HandleError(c, &exception.Unauthorized{Message: "Unauthorized"})
			c.Abort()
			return
		}

		userId, err := helpers.ValidateToken(tokenString)
		if err != nil {
			exception.HandleError(c, &exception.Unauthorized{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", *userId)
		c.Next()
	}
}