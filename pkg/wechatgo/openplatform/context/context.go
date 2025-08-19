package context

import (
	"bookkeeping/pkg/wechatgo/openplatform/config"
)

// Context struct
type Context struct {
	*config.Config
}
