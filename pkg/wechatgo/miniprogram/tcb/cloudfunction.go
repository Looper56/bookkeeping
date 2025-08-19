package tcb

import (
	"bookkeeping/pkg/wechatgo/api"
	"bookkeeping/pkg/wechatgo/util"
)

// InvokeCloudFunctionRes 云函数调用返回结果
type InvokeCloudFunctionRes struct {
	util.CommonError
	RespData string `json:"resp_data"` // 云函数返回的buffer
}

// InvokeCloudFunction 云函数调用
// reference:
// https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/functions/invokeCloudFunction.html
func (tcb *Tcb) InvokeCloudFunction(env, name, args string) (*InvokeCloudFunctionRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := api.InvokeCloudFunctionUri(accessToken, env, name)
	response, err := util.HTTPPost(uri, args)
	if err != nil {
		return nil, err
	}
	invokeCloudFunctionRes := &InvokeCloudFunctionRes{}
	err = util.DecodeWithError(response, invokeCloudFunctionRes, "InvokeCloudFunction")
	return invokeCloudFunctionRes, err
}
