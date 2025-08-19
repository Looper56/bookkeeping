package api

import "fmt"

const (
	officialUserInfoURL = "/cgi-bin/user/info"
	updateRemarkURL     = "/cgi-bin/user/info/updateremark"
	userListURL         = "/cgi-bin/user/get"
)

// OfficialUserInfoUri 获取用户信息接口
func OfficialUserInfoUri(accessToken, openid string) string {
	return fmt.Sprintf("%s%s?access_token=%s&openid=%s&lang=zh_CN",
		GetBaseUrl(), officialUserInfoURL, accessToken, openid)
}

// UpdateRemarkUri 更新用户标签
func UpdateRemarkUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s",
		GetBaseUrl(), updateRemarkURL, accessToken)
}

// UserListUri 获取用户列表
func UserListUri() string {
	return fmt.Sprintf("%s%s",
		GetBaseUrl(), userListURL)
}

const (
	tagCreateURL         = "/cgi-bin/tags/create"
	tagGetURL            = "/cgi-bin/tags/get"
	tagUpdateURL         = "/cgi-bin/tags/update"
	tagDeleteURL         = "/cgi-bin/tags/delete"
	tagUserListURL       = "/cgi-bin/user/tag/get"
	tagBatchTaggingURL   = "/cgi-bin/tags/members/batchtagging"
	tagBatchUntaggingURL = "/cgi-bin/tags/members/batchuntagging"
	tagUserTidListURL    = "/cgi-bin/tags/getidlist"
)

// TagCreateUri tagCreateURL
func TagCreateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), tagCreateURL, accessToken)
}

// TagGetUri tagGetURL
func TagGetUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), tagGetURL, accessToken)
}

// TagUpdateUri tagUpdateURL
func TagUpdateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), tagUpdateURL, accessToken)
}

// TagDeleteUri tagDeleteURL
func TagDeleteUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), tagDeleteURL, accessToken)
}

// TagUserListUri tagUserListURL
func TagUserListUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), tagUserListURL, accessToken)
}

// TagBatchTaggingUri tagBatchTaggingURL
func TagBatchTaggingUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), tagBatchTaggingURL, accessToken)
}

// TagBatchUntaggingUri tagBatchUntaggingURL
func TagBatchUntaggingUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), tagBatchUntaggingURL, accessToken)
}

// TagUserTidListUri tagUserTidListURL
func TagUserTidListUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), tagUserTidListURL, accessToken)
}

const (
	// 获取微信服务器IP地址
	// 文档：https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
	getCallbackIPURL  = "/cgi-bin/getcallbackip"
	getAPIDomainIPURL = "/cgi-bin/get_api_domain_ip"
	// 清理接口调用次数
	clearQuotaURL = "/cgi-bin/clear_quota"
)

// GetCallbackIPUri 获取微信服务器IP地址
func GetCallbackIPUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getCallbackIPURL, accessToken)
}

// GetAPIDomainIPUri 获取微信服务器域名
func GetAPIDomainIPUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getAPIDomainIPURL, accessToken)
}

// ClearQuotaUri 清理接口调用次数
func ClearQuotaUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), clearQuotaURL, accessToken)
}

const (
	qrCreateURL = "/cgi-bin/qrcode/create"
)

// QrCreateUri 二维码
func QrCreateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), qrCreateURL, accessToken)
}

const (
	// 全员或者按用户标准发送消息
	sendURLByTagURL = "/cgi-bin/message/mass/sendall"
	// 按openid列表发送消息
	sendURLByOpenIDURL = "/cgi-bin/message/mass/send"
	// 取消群发消息
	deleteSendURL = "/cgi-bin/message/mass/delete"
)

// SendURLByTagUri 全员或者按用户标准发送消息
func SendURLByTagUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), sendURLByTagURL, accessToken)
}

// SendURLByOpenIDUri 按openid列表发送消息
func SendURLByOpenIDUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), sendURLByOpenIDURL, accessToken)
}

// DeleteSendUri 取消群发消息
func DeleteSendUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), deleteSendURL, accessToken)
}

const (
	customerSendMessageURL = "/cgi-bin/message/custom/send"
)

// CustomerSendMessageUri customerSendMessageURL
func CustomerSendMessageUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), customerSendMessageURL, accessToken)
}
