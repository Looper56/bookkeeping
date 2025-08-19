package message

import (
	"bookkeeping/pkg/wechatgo/api"
	"bookkeeping/pkg/wechatgo/officialaccount/context"
	"bookkeeping/pkg/wechatgo/util"
	"encoding/json"
	"fmt"
)

// Template 模板消息
type Template struct {
	*context.Context
}

// NewTemplate 实例化
func NewTemplate(context *context.Context) *Template {
	tpl := new(Template)
	tpl.Context = context
	return tpl
}

// TemplateMessage 发送的模板消息内容
type TemplateMessage struct {
	// 必须, 接受者OpenID
	ToUser string `json:"touser"`
	// 必须, 模版ID
	TemplateID string `json:"template_id"`
	// 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	URL string `json:"url,omitempty"`
	// 可选, 整个消息的颜色, 可以不设置
	Color string `json:"color,omitempty"`
	// 必须, 模板数据
	Data map[string]*TemplateDataItem `json:"data"`

	MiniProgram struct {
		// 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		AppID string `json:"appid"`
		// 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
		PagePath string `json:"pagepath"`
	} `json:"miniprogram"` // 可选,跳转至小程序地址
}

// TemplateDataItem 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// resTemplateSend 发送返回数据
type resTemplateSend struct {
	util.CommonError

	MsgID int64 `json:"msgid"`
}

// Send 发送模板消息
func (tpl *Template) Send(msg *TemplateMessage) (msgID int64, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := api.SendTemplateUri(accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	var result resTemplateSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	msgID = result.MsgID
	return
}

type resTemplateList struct {
	util.CommonError

	TemplateList []*TemplateItem `json:"template_list"`
}

// TemplateItem 模板消息
type TemplateItem struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

// List 获取模板列表
func (tpl *Template) List() (templateList []*TemplateItem, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := api.GetAllPrivateTemplateUri(accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res resTemplateList
	err = util.DecodeWithError(response, &res, "ListTemplate")
	if err != nil {
		return
	}
	templateList = res.TemplateList
	return
}
