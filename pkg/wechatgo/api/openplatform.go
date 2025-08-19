package api

import "fmt"

const (
	componentAccessTokenURL = "/cgi-bin/component/api_component_token"
	getPreCodeURL           = "/cgi-bin/component/api_create_preauthcode"
	queryAuthURL            = "/cgi-bin/component/api_query_auth"
	refreshTokenURL         = "/cgi-bin/component/api_authorizer_token"
	getComponentInfoURL     = "/cgi-bin/component/api_get_authorizer_info"
	// 获取授权方选项信息
	// getComponentConfigURL = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option
	// ?component_access_token=%s"
	// 获取已授权的账号信息
	// getAuthorizerListURL = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_list
	// ?component_access_token=%s"
)

// ComponentAccessTokenUri component access token接口
func ComponentAccessTokenUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), componentAccessTokenURL)
}

// GetPreCodeUri 授权码获取授权信息接口
func GetPreCodeUri(at string) string {
	return fmt.Sprintf("%s%s?component_access_token=%s", GetBaseUrl(), getPreCodeURL, at)
}

// GetPreCodeUri 授权码获取授权信息接口
func QueryAuthUri(at string) string {
	return fmt.Sprintf("%s%s?component_access_token=%s", GetBaseUrl(), queryAuthURL, at)
}

// RefreshTokenUri 更新access_token接口
func RefreshTokenUri(at string) string {
	return fmt.Sprintf("%s%s?component_access_token=%s", GetBaseUrl(), refreshTokenURL, at)
}

// ComponentInfoUri 获取第三方平台信息接口
func ComponentInfoUri(at string) string {
	return fmt.Sprintf("%s%s?component_access_token=%s", GetBaseUrl(), getComponentInfoURL, at)
}

const (
	platformAccessTokenURL = "/sns/oauth2/component/access_token"
)

// PlatformAccessTokenUri platformAccessTokenURL
func PlatformAccessTokenUri(appID, code, oauthAppID, componentAccessToken string) string {
	return fmt.Sprintf("%s%s?appid=%s&code=%s&grant_type=authorization_code"+
		"&component_appid=%s&component_access_token=%s", GetBaseUrl(),
		platformAccessTokenURL, appID, code, oauthAppID, componentAccessToken)
}
