package base

import (
	"context"

	"github.com/gin-gonic/gin"
)

// GetContext ...
func GetContext(c *gin.Context) context.Context {
	return c.MustGet("ctx").(context.Context)
}

// SetContext ...
func SetContext(c *gin.Context, ctx context.Context) {
	c.Set("ctx", ctx)
}
