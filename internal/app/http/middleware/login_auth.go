package middleware

import (
	"bookkeeping/internal/model"
	"bookkeeping/internal/repository"
	"bookkeeping/internal/service"
	"fmt"

	"github.com/Looper56/plugin/web"
	"github.com/gin-gonic/gin"
)

var isWhiteFields = []string{"is_white"}

// LoginAuth miniProgram auth
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionService := service.NewSessionService()
		userRepository := repository.NewUserRepository()
		token := c.GetHeader("api-token")
		if token == "" {
			cookie, err := c.Cookie("api-token")
			if err != nil {
				web.Failure(c, web.BadParamsErr, gin.H{
					"fields": "token not found",
				})
				return
			}
			token = cookie
		}
		session, err := sessionService.GetSession(c, token)
		if err != nil {
			web.Failure(c, web.NotFoundErr, gin.H{
				"fields": "get session fail",
			})
			return
		}
		if session == nil || session.UserUID == "" {
			web.Failure(c, web.NotAllowed, gin.H{
				"fields": "need login",
			})
			return
		}
		user, err := userRepository.FindOne(c, &repository.FindOneUserCondition{UID: session.UserUID}, isWhiteFields...)
		if err != nil {
			web.Failure(c, web.NotFoundErr, gin.H{
				"fields": "get session fail",
			})
			return
		}
		if user.IsWhite != model.IsWhite {
			web.Failure(c, web.NotAllowed, gin.H{
				"fields": "user is not in white",
			})
			return
		}
		fmt.Printf("before set session: %+v\n", session)
		c.Set("session", session)
		c.Next()
	}
}

// GetSession get session form middleware
func GetSession(c *gin.Context) *model.Session {
	return c.MustGet("session").(*model.Session)
}
