package controller

import (
	"bookkeeping/internal/app/http/request"
	"bookkeeping/pkg/base"
	"bookkeeping/pkg/upload"

	"github.com/gin-gonic/gin"
)

// UploadController ...
type UploadController struct {
	BaseController
}

// NewUploadController init
func NewUploadController() *UploadController {
	return &UploadController{}
}

// Upload upload
func (u *UploadController) Upload(c *gin.Context) {
	ctx := base.GetContext(c)
	var params request.UploadRequest
	if !u.Bind(c, &params) {
		return
	}
	err := upload.ParseFile(c, params)
	if err != nil {
		u.ReportError(ctx, c, err)
		return
	}
	u.Success(c)
}
