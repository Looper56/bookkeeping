package main

import (
	"bookkeeping/config"
	"bookkeeping/pkg/base"
	validate "bookkeeping/pkg/validator"

	"github.com/Looper56/plugin/logger"
	"github.com/Looper56/plugin/mysql"
	"github.com/Looper56/plugin/redis"
	"github.com/Looper56/plugin/sentry"
	"github.com/Looper56/plugin/web"
	sentrygin "github.com/getsentry/sentry-go/gin"
)

func main() {
	// load base config
	conf := config.ParseConfig()
	// new server
	webServer := web.NewWebServer(conf.App.HttpPort, conf.App.Debug)
	// mySQL plugin
	webServer.Use(mysql.NewPlugin(conf.MySQLs))
	// redis plugin
	webServer.Use(redis.NewPlugin(conf.Redis))
	// logger plugin
	webServer.Use(logger.NewPlugin(conf.Loggers))
	// sentry report
	webServer.Use(sentry.NewPlugin(conf.Sentry))
	// set globe middleware
	// panic handle
	webServer.Middleware(web.ErrorHandlerFunc)
	webServer.Middleware(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	// middleware send context
	webServer.Middleware(base.NewContext())
	// record request_id
	webServer.Middleware(base.RequestID()) // gateway
	// weChat init
	config.InitWechat(conf)
	// route config
	webServer.Route(RegisterRoute)
	// register validate
	validate.RegisterValidator()
	// error config
	webServer.SetErrorConfigs(config.ErrorConfig)
	// start server
	webServer.Run()
}
