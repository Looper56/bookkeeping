package main

import (
	"bookkeeping/config"
	"bookkeeping/internal/app/http/controller"
	"bookkeeping/internal/app/http/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoute register route
func RegisterRoute(r *gin.Engine) {
	api := r.Group("api/v1")
	uploadRoute(api)
	loginRoute(api)
	userRoute(api)
	openPlateFormRoute(api)
	officialAccountRoute(api)
	innerAPI := r.Group("inner/v1")
	innerAPIRoute(innerAPI)
	api.GET("/req_test", func(c *gin.Context) {
		c.String(http.StatusOK, config.Config.OA.Token)
	})
	r.LoadHTMLFiles("assets/templates/login_page.html")
}

// wechatRoute function of WeChat
func openPlateFormRoute(r *gin.RouterGroup) {
	api := r.Group("/opf")
	openPlatFromController := controller.NewOpenPlatFormController()
	api.POST("/ticket/callback", openPlatFromController.VerifyTicket)
	api.GET("/auth", openPlatFromController.RedirectAuthPage)
	api.GET("/login_page", openPlatFromController.ComponentLoginPage)
	callbackPath := fmt.Sprintf("%s%s%s", "/", config.Config.WeChat.OfficialAccountAppId, "/callback")
	api.GET(callbackPath, openPlatFromController.ComponentCallback)
	api.GET("/signature", openPlatFromController.JsSDK)
}

// officialAccountRoute ...
func officialAccountRoute(r *gin.RouterGroup) {
	api := r.Group("/oa")
	officialAccountController := controller.NewOfficialAccountController()
	officialAccountAppId := config.Config.WeChat.OfficialAccountAppId
	relativePath := fmt.Sprintf("%s%s%s", "/", officialAccountAppId, "/callback")
	api.POST(relativePath, officialAccountController.Serve)
	api.POST("/mini_app/callback", officialAccountController.Serve)
}

// loginRoute miniProgram auth
func loginRoute(r *gin.RouterGroup) {
	api := r.Group("/login")
	loginController := controller.NewLoginController()
	api.GET("/mp_auth", loginController.MPAuth)
}

// userRoute user manager
func userRoute(r *gin.RouterGroup) {
	api := r.Group("/user")
	userController := controller.NewUserController()
	api.POST("/save", userController.SaveMPUser)
	api.GET("/info", middleware.LoginAuth(), userController.GetUserInfo)
	api.PUT("/info", middleware.LoginAuth(), userController.UpdateUser)
}

// uploadRoute upload file
func uploadRoute(r *gin.RouterGroup) {
	api := r.Group("/upload")
	uploadController := controller.NewUploadController()
	api.POST("/file", uploadController.Upload)
}

// innerAPIRoute ...
func innerAPIRoute(r *gin.RouterGroup) {
}
