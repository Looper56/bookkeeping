package controller

import (
	"bookkeeping/internal/app/http/middleware"
	"bookkeeping/internal/app/http/request"
	"bookkeeping/internal/service"
	"bookkeeping/pkg/base"

	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct {
	BaseController
	userService *service.UserService
}

// NewUserController init
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// SaveMPUser save user info
func (u *UserController) SaveMPUser(c *gin.Context) {
	ctx := base.GetContext(c)
	var params request.SaveUserRequest
	if !u.Bind(c, &params) {
		return
	}
	session := middleware.GetSession(c)
	err := u.userService.MPUserInfo(ctx, session, params.EncryptedData, params.IV)
	if err != nil {
		u.Failure(c, err)
		return
	}
	u.Success(c)
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	ctx := base.GetContext(c)
	session := middleware.GetSession(c)
	user, err := u.userService.UserInfo(ctx, session)
	if err != nil {
		u.Failure(c, err)
		return
	}
	u.Success(c, user)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var params request.UpdateUserRequest
	if !u.Bind(c, &params) {
		return
	}
	ctx := base.GetContext(c)
	session := middleware.GetSession(c)
	err := u.userService.UpdateUser(ctx, session.OpenID, params)
	if err != nil {
		u.Failure(c, err)
		return
	}
	u.Success(c)
}
