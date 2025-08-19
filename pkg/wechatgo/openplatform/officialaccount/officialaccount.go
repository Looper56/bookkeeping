package officialaccount

import (
	"bookkeeping/pkg/wechatgo/credential"
	"bookkeeping/pkg/wechatgo/officialaccount"
	offConfig "bookkeeping/pkg/wechatgo/officialaccount/config"
	"bookkeeping/pkg/wechatgo/officialaccount/js"
	"bookkeeping/pkg/wechatgo/officialaccount/oauth"
	opContext "bookkeeping/pkg/wechatgo/openplatform/context"
)

// OfficialAccount 代公众号实现业务
type OfficialAccount struct {
	// 授权的公众号的appID
	appID string
	*officialaccount.OfficialAccount
}

// NewOfficialAccount 实例化
// appID :为授权方公众号 APPID，非开放平台第三方平台 APPID
func NewOfficialAccount(opCtx *opContext.Context, appID string) *OfficialAccount {
	officialAccount := officialaccount.NewOfficialAccount(&offConfig.Config{
		AppID:          opCtx.AppID,
		EncodingAESKey: opCtx.EncodingAESKey,
		Token:          opCtx.Token,
		Cache:          opCtx.Cache,
	})
	// 设置获取access_token的函数
	officialAccount.SetAccessTokenHandle(NewDefaultAuthrAccessToken(opCtx, appID))
	// 设置获取verify_ticket的函数
	officialAccount.SetVerifyTicketHandle(NewDefaultVerifyTicket(opCtx))
	return &OfficialAccount{appID: appID, OfficialAccount: officialAccount}
}

// PlatformOauth 平台代发起oauth2网页授权
func (officialAccount *OfficialAccount) PlatformOauth() *oauth.Oauth {
	return oauth.NewOauth(officialAccount.GetContext())
}

// PlatformJs 平台代获取js-sdk配置
func (officialAccount *OfficialAccount) PlatformJs() *js.Js {
	return js.NewJs(officialAccount.GetContext())
}

// DefaultAuthrAccessToken 默认获取授权ak的方法
type DefaultAuthrAccessToken struct {
	opCtx *opContext.Context
	appID string
}

// NewDefaultAuthrAccessToken New
func NewDefaultAuthrAccessToken(opCtx *opContext.Context, appID string) credential.AccessTokenHandle {
	return &DefaultAuthrAccessToken{
		opCtx: opCtx,
		appID: appID,
	}
}

// GetAccessToken 获取ak
func (ak *DefaultAuthrAccessToken) GetAccessToken() (string, error) {
	return ak.opCtx.GetAuthrAccessToken(ak.appID)
}

// DefaultVerifyTicket 默认获取verify ticket的方法
type DefaultVerifyTicket struct {
	opCtx *opContext.Context
}

// NewDefaultVerifyTicket new
func NewDefaultVerifyTicket(opCtx *opContext.Context) credential.VerifyTicketHandle {
	return &DefaultVerifyTicket{
		opCtx: opCtx,
	}
}

// SetVerifyTicket 获取verify ticket
func (vt *DefaultVerifyTicket) SetVerifyTicket(verifyTicket string) error {
	// 更新component access token，保证一直有值
	_, err := vt.opCtx.SetComponentAccessToken(verifyTicket)
	if err != nil {
		return err
	}

	return nil
}
