package base

import (
	"github.com/Looper56/plugin/logger"
	"github.com/gin-gonic/gin"
)

// RequestID record request ID
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("uid")
		if reqID == "" {
			reqID = c.GetHeader("name")
		}
		ctx := GetContext(c)
		logger.WithContext(ctx, "request_id", reqID)
		SetContext(c, ctx)
		c.Next()
	}
}
