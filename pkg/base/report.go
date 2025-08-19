package base

import (
	"context"

	"github.com/Looper56/plugin/logger"
	"github.com/Looper56/plugin/web"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// ReportError ...
func ReportError(ctx context.Context, c *gin.Context, err error) {
	logger.ErrorContext(ctx, err.Error())
	web.Failure(c, err)
	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			hub.CaptureException(err)
		})
	}
}
