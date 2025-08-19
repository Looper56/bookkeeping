package api

var (
	defaultBaseUrl = "https://api.weixin.qq.com"
)

// GetBaseUrl 获取微信接口域名
func GetBaseUrl() string {
	return defaultBaseUrl
}

// SetBaseUrl 设置微信接口域名
func SetBaseUrl(baseUrl string) {
	defaultBaseUrl = baseUrl
}
