package api

import "fmt"

const (
	code2SessionURL     = "/sns/jscode2session"
	customerSendMessage = "/cgi-bin/message/custom/send"
)

// Code2SessionUri 获取用户access_token接口
func Code2SessionUri(appID, appSecret, code string) string {
	return fmt.Sprintf("%s%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		GetBaseUrl(), code2SessionURL, appID, appSecret, code)
}

// MPCustomerSendMessageUri 获取小程序模板消息接口
func MPCustomerSendMessageUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), customerSendMessage, accessToken)
}

const (
	createWXAQRCodeURL     = "/cgi-bin/wxaapp/createwxaqrcode?access_token=%s"
	getWXACodeURL          = "/wxa/getwxacode?access_token=%s"
	getWXACodeUnlimitedURL = "/wxa/getwxacodeunlimit?access_token=%s"
)

// CreateWXAQRCodeUri 获取生成QR二维码接口
func CreateWXAQRCodeUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), createWXAQRCodeURL)
}

// GetWXACodeUri 获取生成二维码接口
func GetWXACodeUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getWXACodeURL)
}

// GetWXACodeUnlimitedUri 获取二维码接口
func GetWXACodeUnlimitedUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getWXACodeUnlimitedURL)
}

const (
	// 发送订阅消息
	// developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
	subscribeSendURL = "/cgi-bin/message/subscribe/send"

	// 获取当前帐号下的个人模板列表
	getTemplateURL = "/wxaapi/newtmpl/gettemplate"
)

// SubscribeSendUri 发送订阅消息接口
func SubscribeSendUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), subscribeSendURL, accessToken)
}

// GetTemplateUri 获取当前帐号下的个人模板列表接口
func GetTemplateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getTemplateURL, accessToken)
}

const (
	// 获取用户访问小程序日留存
	getAnalysisDailyRetainURL = "/datacube/getweanalysisappiddailyretaininfo?access_token=%s"
	// 获取用户访问小程序月留存
	getAnalysisMonthlyRetainURL = "/datacube/getweanalysisappidmonthlyretaininfo?access_token=%s"
	// 获取用户访问小程序周留存
	getAnalysisWeeklyRetainURL = "/datacube/getweanalysisappidweeklyretaininfo?access_token=%s"
	// 获取用户访问小程序数据概况
	getAnalysisDailySummaryURL = "/datacube/getweanalysisappiddailysummarytrend?access_token=%s"
	// 获取用户访问小程序数据日趋势
	getAnalysisDailyVisitTrendURL = "/datacube/getweanalysisappiddailyvisittrend?access_token=%s"
	// 获取用户访问小程序数据月趋势
	getAnalysisMonthlyVisitTrendURL = "/datacube/getweanalysisappidmonthlyvisittrend?access_token=%s"
	// 获取用户访问小程序数据周趋势
	getAnalysisWeeklyVisitTrendURL = "/datacube/getweanalysisappidweeklyvisittrend?access_token=%s"
	// 获取小程序新增或活跃用户的画像分布数据
	getAnalysisUserPortraitURL = "/datacube/getweanalysisappiduserportrait?access_token=%s"
	// 获取用户小程序访问分布数据
	getAnalysisVisitDistributionURL = "/datacube/getweanalysisappidvisitdistribution?access_token=%s"
	// 访问页面
	getAnalysisVisitPageURL = "/datacube/getweanalysisappidvisitpage?access_token=%s"
)

// GetAnalysisDailyRetainUri 获取用户访问小程序日留存接口
func GetAnalysisDailyRetainUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisDailyRetainURL)
}

// GetAnalysisMonthlyRetainUri 获取用户访问小程序月留存
func GetAnalysisMonthlyRetainUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisMonthlyRetainURL)
}

// GetAnalysisWeeklyRetainUri 获取用户访问小程序周留存
func GetAnalysisWeeklyRetainUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisWeeklyRetainURL)
}

// GetAnalysisDailySummaryUri 获取用户访问小程序周留存
func GetAnalysisDailySummaryUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisDailySummaryURL)
}

// GetAnalysisDailyVisitTrendUri 获取用户访问小程序数据日趋势
func GetAnalysisDailyVisitTrendUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisDailyVisitTrendURL)
}

// GetAnalysisMonthlyVisitTrendUri 获取用户访问小程序数据月趋势
func GetAnalysisMonthlyVisitTrendUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisMonthlyVisitTrendURL)
}

// GetAnalysisWeeklyVisitTrendUri 获取用户访问小程序数据周趋势
func GetAnalysisWeeklyVisitTrendUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisWeeklyVisitTrendURL)
}

// GetAnalysisUserPortraitUri 获取小程序新增或活跃用户的画像分布数据
func GetAnalysisUserPortraitUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisUserPortraitURL)
}

// GetAnalysisVisitDistributionUri 获取用户小程序访问分布数据
func GetAnalysisVisitDistributionUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisVisitDistributionURL)
}

// GetAnalysisVisitPageUri 访问页面
func GetAnalysisVisitPageUri() string {
	return fmt.Sprintf("%s%s", GetBaseUrl(), getAnalysisVisitPageURL)
}

const (
	getAccountBasicInfoURL = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo"
)

// GetAccountBasicInfoUri getAccountBasicInfoURL
func GetAccountBasicInfoUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), getAccountBasicInfoURL, accessToken)
}

const (
	fastregisterweappURL = "https://api.weixin.qq.com/cgi-bin/component/fastregisterweapp"
)

// FastregisterweappUri fastregisterweappURL
func FastregisterweappUri(accessToken string) string {
	return fmt.Sprintf("%s%s?action=create&component_access_token=%s", GetBaseUrl(),
		fastregisterweappURL, accessToken)
}
