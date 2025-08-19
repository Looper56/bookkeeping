package api

import "fmt"

const (
	templateSendURL = "/cgi-bin/message/template/send"
	templateListURL = "/cgi-bin/template/get_all_private_template"
)

// SendTemplateUri 模板消息发送接口
func SendTemplateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), templateSendURL, accessToken)
}

// GetAllPrivateTemplateUri 获取所有模板消息接口
func GetAllPrivateTemplateUri(accessToken string) string {
	return fmt.Sprintf("%s%s?access_token=%s", GetBaseUrl(), templateListURL, accessToken)
}
