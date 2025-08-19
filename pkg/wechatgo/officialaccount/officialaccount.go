package officialaccount

import (
	"bookkeeping/pkg/wechatgo/credential"
	"bookkeeping/pkg/wechatgo/officialaccount/basic"
	"bookkeeping/pkg/wechatgo/officialaccount/broadcast"
	"bookkeeping/pkg/wechatgo/officialaccount/config"
	"bookkeeping/pkg/wechatgo/officialaccount/context"
	"bookkeeping/pkg/wechatgo/officialaccount/datacube"
	"bookkeeping/pkg/wechatgo/officialaccount/device"
	"bookkeeping/pkg/wechatgo/officialaccount/js"
	"bookkeeping/pkg/wechatgo/officialaccount/material"
	"bookkeeping/pkg/wechatgo/officialaccount/menu"
	"bookkeeping/pkg/wechatgo/officialaccount/message"
	"bookkeeping/pkg/wechatgo/officialaccount/oauth"
	"bookkeeping/pkg/wechatgo/officialaccount/server"
	"bookkeeping/pkg/wechatgo/officialaccount/user"
	"net/http"
)

// OfficialAccount 微信公众号相关API
type OfficialAccount struct {
	ctx *context.Context
}

// NewOfficialAccount 实例化公众号API
func NewOfficialAccount(cfg *config.Config) *OfficialAccount {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &OfficialAccount{ctx: ctx}
}

// StartRefreshAccessTokenTicker 启动定时更新access_token，由调用方触发
func (officialAccount *OfficialAccount) StartRefreshAccessTokenTicker() {
	officialAccount.ctx.StartRefreshAccessTokenTicker()
}

// SetAccessTokenHandle 自定义access_token获取方式
func (officialAccount *OfficialAccount) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	officialAccount.ctx.AccessTokenHandle = accessTokenHandle
}

// SetVerifyTicketHandle 自定义verify ticket设置方式
func (officialAccount *OfficialAccount) SetVerifyTicketHandle(verifyTicketHandle credential.VerifyTicketHandle) {
	officialAccount.ctx.VerifyTicketHandle = verifyTicketHandle
}

// GetContext get Context
func (officialAccount *OfficialAccount) GetContext() *context.Context {
	return officialAccount.ctx
}

// GetBasic qr/url 相关配置
func (officialAccount *OfficialAccount) GetBasic() *basic.Basic {
	return basic.NewBasic(officialAccount.ctx)
}

// GetMenu 菜单管理接口
func (officialAccount *OfficialAccount) GetMenu() *menu.Menu {
	return menu.NewMenu(officialAccount.ctx)
}

// GetServer 消息管理：接收事件，被动回复消息管理
func (officialAccount *OfficialAccount) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	srv := server.NewServer(officialAccount.ctx)
	srv.Request = req
	srv.Writer = writer
	return srv
}

// GetAccessToken 获取access_token
func (officialAccount *OfficialAccount) GetAccessToken() (string, error) {
	return officialAccount.ctx.GetAccessToken()
}

// GetOauth oauth2网页授权
func (officialAccount *OfficialAccount) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(officialAccount.ctx)
}

// GetMaterial 素材管理
func (officialAccount *OfficialAccount) GetMaterial() *material.Material {
	return material.NewMaterial(officialAccount.ctx)
}

// GetJs js-sdk配置
func (officialAccount *OfficialAccount) GetJs() *js.Js {
	return js.NewJs(officialAccount.ctx)
}

// GetUser 用户管理接口
func (officialAccount *OfficialAccount) GetUser() *user.User {
	return user.NewUser(officialAccount.ctx)
}

// GetTemplate 模板消息接口
func (officialAccount *OfficialAccount) GetTemplate() *message.Template {
	return message.NewTemplate(officialAccount.ctx)
}

// GetCustomerMessageManager 客服消息接口
func (officialAccount *OfficialAccount) GetCustomerMessageManager() *message.Manager {
	return message.NewMessageManager(officialAccount.ctx)
}

// GetDevice 获取智能设备的实例
func (officialAccount *OfficialAccount) GetDevice() *device.Device {
	return device.NewDevice(officialAccount.ctx)
}

// GetBroadcast 群发消息
func (officialAccount *OfficialAccount) GetBroadcast() *broadcast.Broadcast {
	return broadcast.NewBroadcast(officialAccount.ctx)
}

// GetDataCube 数据统计
func (officialAccount *OfficialAccount) GetDataCube() *datacube.DataCube {
	return datacube.NewCube(officialAccount.ctx)
}
