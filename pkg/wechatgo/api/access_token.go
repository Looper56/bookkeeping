package api

import "fmt"

const (
	// AccessTokenURL 获取access_token的接口
	AccessTokenURL = "/cgi-bin/token"
)

// AccessTokenUri 模板消息发送接口
func AccessTokenUri(appID string, appSecret string) string {
	return fmt.Sprintf("%s%s?grant_type=client_credential&appid=%s&secret=%s",
		GetBaseUrl(), AccessTokenURL, appID, appSecret)
}
