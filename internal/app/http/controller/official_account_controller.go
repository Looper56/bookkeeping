package controller

import (
	"bookkeeping/internal/service"
	"bookkeeping/pkg/wechatgo/officialaccount/message"
	"bookkeeping/pkg/wechatgo/openplatform/officialaccount"
	"fmt"

	"github.com/Looper56/plugin/logger"
	"github.com/gin-gonic/gin"
)

// OfficialAccountController official account ...
type OfficialAccountController struct {
	BaseController
	openOfficeAccount *officialaccount.OfficialAccount
}

// NewOfficialAccountController init
func NewOfficialAccountController() *OfficialAccountController {
	var openOfficeAccount *officialaccount.OfficialAccount
	return &OfficialAccountController{

		openOfficeAccount: openOfficeAccount,
	}
}

// Serve ...
func (o *OfficialAccountController) Serve(c *gin.Context) {
	server := o.openOfficeAccount.GetServer(c.Request, c.Writer)

	officialAccountSvc := service.NewOfficialAccountService(server)

	officialAccountSvc.OnEvent(message.EventSubscribe, officialAccountSvc.UserSubscribe)
	officialAccountSvc.OnEvent(message.EventUnsubscribe, officialAccountSvc.UserUnSubscribe)
	officialAccountSvc.OnMsg(message.MsgTypeText, officialAccountSvc.AutoReply)
	err := officialAccountSvc.Serve()
	if err != nil {
		logger.Error(fmt.Sprintf("official account serve fail: %+v", err))
	}
}
