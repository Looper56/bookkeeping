package api

import "fmt"

const (
	accessTokenURL        = "/sns/oauth2/access_token"
	refreshAccessTokenURL = "/sns/oauth2/refresh_token"
	userInfoURL           = "/sns/userinfo"
	checkAccessTokenURL   = "/sns/auth"
)

// OauthAccessTokenUri 获取用户access_token接口
func OauthAccessTokenUri(appID, appSecret, code string) string {
	return fmt.Sprintf("%s%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		GetBaseUrl(), accessTokenURL, appID, appSecret, code)
}

// OauthRefreshAccessTokenUri 刷新用户access_token接口
func OauthRefreshAccessTokenUri(appID, refreshToken string) string {
	return fmt.Sprintf("%s%s?appid=%s&grant_type=refresh_token&refresh_token=%s",
		GetBaseUrl(), refreshAccessTokenURL, appID, refreshToken)
}

// OauthUserInfoUri 获取用户信息接口
func OauthUserInfoUri(accessToken, openid string) string {
	return fmt.Sprintf("%s%s?access_token=%s&openid=%s&lang=zh_CN",
		GetBaseUrl(), userInfoURL, accessToken, openid)
}

// OauthCheckUserInfoUri 检查用户信息接口
func OauthCheckUserInfoUri(accessToken, openid string) string {
	return fmt.Sprintf("%s%s?access_token=%s&openid=%s",
		GetBaseUrl(), checkAccessTokenURL, accessToken, openid)
}
