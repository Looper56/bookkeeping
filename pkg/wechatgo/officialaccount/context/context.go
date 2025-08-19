package context

import (
	"bookkeeping/pkg/wechatgo/credential"
	"bookkeeping/pkg/wechatgo/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
	credential.VerifyTicketHandle
}
