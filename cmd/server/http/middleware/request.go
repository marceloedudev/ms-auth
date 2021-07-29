package middleware

import (
	"ms-auth/pkg/uuid"

	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewUUIDString())
		c.Next()
	}
}
