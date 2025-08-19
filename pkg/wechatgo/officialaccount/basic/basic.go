package basic

import (
	"bookkeeping/pkg/wechatgo/api"
	"bookkeeping/pkg/wechatgo/officialaccount/context"
	"bookkeeping/pkg/wechatgo/util"
)

// Basic struct
type Basic struct {
	*context.Context
}

// NewBasic 实例
func NewBasic(context *context.Context) *Basic {
	basic := new(Basic)
	basic.Context = context
	return basic
}

// IPListRes 获取微信服务器IP地址 返回结果
type IPListRes struct {
	util.CommonError
	IPList []string `json:"ip_list"`
}

// GetCallbackIP 获取微信callback IP地址
func (basic *Basic) GetCallbackIP() ([]string, error) {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := api.GetCallbackIPUri(ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetCallbackIP")
	return ipListRes.IPList, err
}

// GetAPIDomainIP 获取微信API接口 IP地址
func (basic *Basic) GetAPIDomainIP() ([]string, error) {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := api.GetAPIDomainIPUri(ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetAPIDomainIP")
	return ipListRes.IPList, err
}

// ClearQuota 清理接口调用次数
func (basic *Basic) ClearQuota() error {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return err
	}
	url := api.ClearQuotaUri(ak)
	data, err := util.PostJSON(url, map[string]string{
		"appid": basic.AppID,
	})
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(data, "ClearQuota")
}
