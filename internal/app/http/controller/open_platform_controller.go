package controller

import (
	"bookkeeping/config"
	"bookkeeping/internal/service"
	"bookkeeping/pkg/wechatgo/openplatform"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OpenPlatFormController OpenPlatForm controller
type OpenPlatFormController struct {
	BaseController
	openPlatFormService *service.OpenPlatFormService
	openPlatForm        *openplatform.OpenPlatform
}

// NewOpenPlatFormController init
func NewOpenPlatFormController() *OpenPlatFormController {
	var openPlatForm *openplatform.OpenPlatform
	return &OpenPlatFormController{
		openPlatFormService: service.NewOpenPlatFormService(),
		openPlatForm:        openPlatForm,
	}
}

// VerifyTicket verify ticket
func (o *OpenPlatFormController) VerifyTicket(c *gin.Context) {
	server := o.openPlatForm.GetServer(c.Request, c.Writer)
	server.SkipValidate(config.Config.WeChat.Debug)
	err := o.openPlatFormService.Serve(server)
	if err != nil {
		o.Failure(c, err)
		return
	}
	o.Success(c)
}

// RedirectAuthPage redirect to auth page by api route
func (o *OpenPlatFormController) RedirectAuthPage(c *gin.Context) {
	uri := fmt.Sprintf("%s%s", config.Config.WeChat.AppURL, "/api/v1/wechat/login_page")
	c.HTML(http.StatusOK, "login_page.html", gin.H{
		"uri": uri,
	})
}

// ComponentLoginPage generate auth page QR code
func (o *OpenPlatFormController) ComponentLoginPage(c *gin.Context) {
	redirectURL := fmt.Sprintf("%s", "callback")
	uri, err := o.openPlatForm.GetComponentLoginPage(redirectURL, 3, config.Config.WeChat.OfficialAccountAppId)
	if err != nil {
		o.Failure(c, err)
		return
	}
	c.Redirect(http.StatusFound, uri)
}

// ComponentCallback third account empower callback
func (o *OpenPlatFormController) ComponentCallback(c *gin.Context) {
	authCode := c.Query("auth_code")
	_, err := o.openPlatForm.QueryAuthCode(authCode)
	if err != nil {
		o.Failure(c, err)
		return
	}
	o.Success(c)
}

// JsSDK third account empower callback address TODO: need find content in doc
func (o *OpenPlatFormController) JsSDK(c *gin.Context) {
	jsURL := c.Query("url")
	officeAccount := o.openPlatForm.GetOfficialAccount(config.Config.WeChat.OfficialAccountAppId)
	conf, err := officeAccount.GetJs().GetConfig(jsURL)
	if err != nil {
		o.Failure(c, err)
		return
	}
	conf.JsApiList = []string{""} // updateAppMessageShareData ...
	o.Success(c, conf)
}
