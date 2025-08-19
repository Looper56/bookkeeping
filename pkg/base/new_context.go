package base

import (
	"context"

	"github.com/gin-gonic/gin"
)

// NewContext set current request context
func NewContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		SetContext(c, context.Background())
		c.Next()
	}
}
