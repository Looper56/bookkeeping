package openplatform

import (
	"bookkeeping/pkg/wechatgo/officialaccount/server"
	"bookkeeping/pkg/wechatgo/openplatform/account"
	"bookkeeping/pkg/wechatgo/openplatform/config"
	"bookkeeping/pkg/wechatgo/openplatform/context"
	"bookkeeping/pkg/wechatgo/openplatform/miniprogram"
	"bookkeeping/pkg/wechatgo/openplatform/officialaccount"
	"net/http"
)

// OpenPlatform 微信开放平台相关api
type OpenPlatform struct {
	*context.Context
}

// NewOpenPlatform new open platform
func NewOpenPlatform(cfg *config.Config) *OpenPlatform {
	if cfg.Cache == nil {
		panic("cache 未设置")
	}
	ctx := &context.Context{
		Config: cfg,
	}
	return &OpenPlatform{ctx}
}

// GetServer get server
func (openPlatform *OpenPlatform) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	off := officialaccount.NewOfficialAccount(openPlatform.Context, "")
	return off.GetServer(req, writer)
}

// GetOfficialAccount 公众号代处理
func (openPlatform *OpenPlatform) GetOfficialAccount(appID string) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(openPlatform.Context, appID)
}

// GetMiniProgram 小程序代理
func (openPlatform *OpenPlatform) GetMiniProgram(appID string) *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(openPlatform.Context, appID)
}

// GetAccountManager 账号管理
func (openPlatform *OpenPlatform) GetAccountManager() *account.Account {
	return account.NewAccount(openPlatform.Context)
}
