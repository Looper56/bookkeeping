package context

import (
	"bookkeeping/pkg/wechatgo/credential"
	"bookkeeping/pkg/wechatgo/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
