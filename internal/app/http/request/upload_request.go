package request

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

// UploadRequest upload request
type UploadRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// Validate ...
func (UploadRequest) Validate(c *gin.Context) error {
	return nil
}
