package request

import "github.com/gin-gonic/gin"

// PaginationRequest paginate request
type PaginationRequest struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

// Validate ...
func (PaginationRequest) Validate(c *gin.Context) error {
	return nil
}
