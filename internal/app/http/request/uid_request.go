package request

import "github.com/gin-gonic/gin"

// UIDRequest uid request
type UIDRequest struct {
	UID string `json:"uid" binding:"required"`
}

// Validate ...
func (UIDRequest) Validate(c *gin.Context) error {
	return nil
}
