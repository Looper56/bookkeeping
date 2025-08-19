package controller

import (
	"context"

	"github.com/Looper56/plugin/web"

	"bookkeeping/pkg/base"

	"github.com/gin-gonic/gin"
)

// BaseController ...
type BaseController struct {
	web.BaseController
}

// ReportError error report
func (b *BaseController) ReportError(ctx context.Context, c *gin.Context, err error) {
	base.ReportError(ctx, c, err)
}
