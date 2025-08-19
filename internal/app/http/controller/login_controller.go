package controller

import (
	"bookkeeping/internal/app/http/request"
	"bookkeeping/internal/service"
	"bookkeeping/pkg/base"

	"github.com/gin-gonic/gin"
)

// LoginController login controller
type LoginController struct {
	BaseController
	loginService *service.LoginService
}

// NewLoginController init controller
func NewLoginController() *LoginController {
	return &LoginController{
		loginService: service.NewLoginService(),
	}
}

// MPAuth miniProgram auth
func (l *LoginController) MPAuth(c *gin.Context) {
	ctx := base.GetContext(c)
	var params request.MPLoginRequest
	if !l.Bind(c, &params) {
		return
	}
	loginState, err := l.loginService.MPAuth(ctx, params.Code)
	if err != nil {
		l.Failure(c, err)
		return
	}
	l.Success(c, loginState)
}
