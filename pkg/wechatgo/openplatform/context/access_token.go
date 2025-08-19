// Package context 开放平台相关context
// verify ticket 在redis中存储key为：verify_ticket_{component_appid}
// Component access token在redis中的存储key为：component_access_token_{component_appid}
// Authorizer access token在redis中存储key为：authorizer_access_token_{authorizer_appid}
package context

import (
	"bookkeeping/pkg/wechatgo/api"
	"bookkeeping/pkg/wechatgo/util"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	componentLoginURL = "https://mp.weixin.qq.com/cgi-bin/componentloginpage" +
		"?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d&biz_appid=%s"
	bindComponentURL = "https://mp.weixin.qq.com/safe/bindcomponent" +
		"?action=bindcomponent&auth_type=%d&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&biz_appid=%s" +
		"#wechat_redirect"
)

// ComponentAccessToken 第三方平台
type ComponentAccessToken struct {
	AccessToken string `json:"component_access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// componentAccessTokenStore access token存储
type componentAccessTokenStore struct {
	Token     *ComponentAccessToken
	ExpiresAt time.Time
}

// GetComponentAccessToken 获取 ComponentAccessToken
func (ctx *Context) GetComponentAccessToken() (string, error) {
	accessTokenCacheKey := fmt.Sprintf("component_access_token_%s", ctx.AppID)
	val := ctx.Cache.Get(accessTokenCacheKey)
	if val == nil {
		return "", fmt.Errorf("cann't get component access token")
	}

	atStore := &componentAccessTokenStore{}
	keyStr := val.(string)
	err := json.Unmarshal([]byte(keyStr), atStore)
	if err != nil {
		return "", fmt.Errorf("unmarshal component access token fail")
	}

	return atStore.Token.AccessToken, nil
}

// SetComponentAccessToken 通过 component_verify_ticket 获取 ComponentAccessToken
func (ctx *Context) SetComponentAccessToken(verifyTicket string) (*ComponentAccessToken, error) {
	err := ctx.saveVerifyTicket(verifyTicket)
	if err != nil {
		return nil, err
	}

	at := ctx.getComponentAccessTokenNotExpired()
	if at != nil {
		return at, nil
	}
	at, err = ctx.updateComponentAccessToken(verifyTicket)
	if err != nil {
		return nil, err
	}

	return at, nil
}

// GetPreCode 获取预授权码
func (ctx *Context) GetPreCode() (string, error) {
	cat, err := ctx.GetComponentAccessToken()
	if err != nil {
		return "", err
	}
	req := map[string]string{
		"component_appid": ctx.AppID,
	}
	uri := api.GetPreCodeUri(cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return "", err
	}

	var ret struct {
		PreCode string `json:"pre_auth_code"`
	}
	if err := json.Unmarshal(body, &ret); err != nil {
		return "", err
	}

	return ret.PreCode, nil
}

// GetComponentLoginPage 获取第三方公众号授权链接(扫码授权)
func (ctx *Context) GetComponentLoginPage(redirectURI string, authType int, bizAppID string) (string, error) {
	code, err := ctx.GetPreCode()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(componentLoginURL, ctx.AppID, code, url.QueryEscape(redirectURI), authType, bizAppID), nil
}

// GetBindComponentURL 获取第三方公众号授权链接(链接跳转，适用移动端)
func (ctx *Context) GetBindComponentURL(redirectURI string, authType int, bizAppID string) (string, error) {
	code, err := ctx.GetPreCode()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(bindComponentURL, authType, ctx.AppID, code, url.QueryEscape(redirectURI), bizAppID), nil
}

// ID 微信返回接口中各种类型字段
type ID struct {
	ID int `json:"id"`
}

// AuthBaseInfo 授权的基本信息
type AuthBaseInfo struct {
	AuthrAccessToken
	FuncInfo []AuthFuncInfo `json:"func_info"`
}

// AuthFuncInfo 授权的接口内容
type AuthFuncInfo struct {
	FuncscopeCategory ID `json:"funcscope_category"`
}

// AuthrAccessToken 授权方AccessToken
type AuthrAccessToken struct {
	Appid        string `json:"authorizer_appid"`
	AccessToken  string `json:"authorizer_access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"authorizer_refresh_token"`
}

// QueryAuthCode 使用授权码换取公众号或小程序的接口调用凭据和授权信息
func (ctx *Context) QueryAuthCode(authCode string) (*AuthBaseInfo, error) {
	cat, err := ctx.GetComponentAccessToken()
	if err != nil {
		return nil, err
	}

	req := map[string]string{
		"component_appid":    ctx.AppID,
		"authorization_code": authCode,
	}
	uri := api.QueryAuthUri(cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	var ret struct {
		util.CommonError
		Info *AuthBaseInfo `json:"authorization_info"`
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("QueryAuthCode error : errcode=%v , errmsg=%v", ret.ErrCode, ret.ErrMsg)
		return nil, err
	}

	// 保存信息到db中
	err = ctx.SaveAuthInfo(ret.Info)
	if err != nil {
		return nil, err
	}

	return ret.Info, nil
}

// RefreshAuthrToken 获取（刷新）授权公众号或小程序的接口调用凭据（令牌）
func (ctx *Context) RefreshAuthrToken(appID, refreshToken string) (*AuthrAccessToken, error) {
	cat, err := ctx.GetComponentAccessToken()
	if err != nil {
		return nil, err
	}
	req := map[string]string{
		"component_appid":          ctx.AppID,
		"authorizer_appid":         appID,
		"authorizer_refresh_token": refreshToken,
	}
	uri := api.RefreshTokenUri(cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	aat := &AuthrAccessToken{}
	if err := json.Unmarshal(body, aat); err != nil {
		return nil, err
	}
	if aat.AccessToken == "" {
		err = fmt.Errorf("get authorizer access token fail: %s", string(body))
		log.Error(err.Error())
		return nil, err
	}
	aat.Appid = appID

	err = ctx.UpdateAuthrAccessToken(aat)
	if err != nil {
		return nil, err
	}
	err = ctx.CacheAuthrAccessToken(appID, aat.AccessToken, aat.ExpiresIn)
	if err != nil {
		return nil, err
	}
	return aat, nil
}

// GetAuthrAccessToken 获取授权方AccessToken
func (ctx *Context) GetAuthrAccessToken(appid string) (string, error) {
	authrTokenKey := "authorizer_access_token_" + appid
	val := ctx.Cache.Get(authrTokenKey)
	if val == nil {
		return "", fmt.Errorf("cannot get authorizer %s access token", appid)
	}
	return val.(string), nil
}

// AuthorizerInfo 授权方详细信息
type AuthorizerInfo struct {
	NickName        string `json:"nick_name"`
	HeadImg         string `json:"head_img"`
	ServiceTypeInfo ID     `json:"service_type_info"`
	VerifyTypeInfo  ID     `json:"verify_type_info"`
	UserName        string `json:"user_name"`
	PrincipalName   string `json:"principal_name"`
	BusinessInfo    struct {
		OpenStore string `json:"open_store"`
		OpenScan  string `json:"open_scan"`
		OpenPay   string `json:"open_pay"`
		OpenCard  string `json:"open_card"`
		OpenShake string `json:"open_shake"`
	}
	Alias     string `json:"alias"`
	QrcodeURL string `json:"qrcode_url"`
}

// GetAuthrInfo 获取授权方的帐号基本信息
func (ctx *Context) GetAuthrInfo(appid string) (*AuthorizerInfo, *AuthBaseInfo, error) {
	cat, err := ctx.GetComponentAccessToken()
	if err != nil {
		return nil, nil, err
	}

	req := map[string]string{
		"component_appid":  ctx.AppID,
		"authorizer_appid": appid,
	}

	uri := api.ComponentInfoUri(cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, nil, err
	}

	var ret struct {
		AuthorizerInfo    *AuthorizerInfo `json:"authorizer_info"`
		AuthorizationInfo *AuthBaseInfo   `json:"authorization_info"`
	}
	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, nil, err
	}

	return ret.AuthorizerInfo, ret.AuthorizationInfo, nil
}
