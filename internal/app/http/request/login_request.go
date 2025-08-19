package request

import "github.com/gin-gonic/gin"

// MPLoginRequest miniProgram auth request
type MPLoginRequest struct {
	Code string `form:"code" binding:"required"`
}

// Validate ...
func (MPLoginRequest) Validate(c *gin.Context) error {
	return nil
}
