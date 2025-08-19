package config

import (
	"bookkeeping/pkg/wechatgo"
	"bookkeeping/pkg/wechatgo/cache"
	"bookkeeping/pkg/wechatgo/miniprogram"
	miniConfig "bookkeeping/pkg/wechatgo/miniprogram/config"
	"bookkeeping/pkg/wechatgo/officialaccount"
	"bookkeeping/pkg/wechatgo/openplatform"
	openOfficialAccount "bookkeeping/pkg/wechatgo/openplatform/officialaccount"
)

// OpenPlatform Operation class...
var OpenPlatform *openplatform.OpenPlatform

// OpenOfficialAccount Operation class...
var OpenOfficialAccount *openOfficialAccount.OfficialAccount

// MiniProgramAccount Operation class...
var MiniProgramAccount *officialaccount.OfficialAccount

// MiniProgram Operation class...
var MiniProgram *miniprogram.MiniProgram

func InitWechat(config *GlobalConfig) {
	wc := wechatgo.NewWechat()

	redisOpts := &cache.RedisOpts{
		// use redis server list default
		Host:        config.Redis[0].Server,
		Password:    config.Redis[0].Password,
		Database:    config.Redis[0].Database,
		MaxActive:   config.Redis[0].MaxConnAgeSeconds,
		MaxIdle:     config.Redis[0].MinIdleConns,
		IdleTimeout: 10000,
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)

	//db, _ := mysql.GetMySQLConnection("default")

	// open platform
	//openConf := &openConfig.Config{
	//	AppID:          config.WeChat.OpenPlatformAppId,
	//	AppSecret:      config.WeChat.OpenPlatformAppSecret,
	//	Token:          config.WeChat.OpenPlatformToken,
	//	EncodingAESKey: config.WeChat.OpenPlatformAesKey,
	//	Cache:          redisCache,
	//	DB:             db,
	//}
	//OpenPlatform = wc.GetOpenPlatform(openConf)

	// start cron update ticket (local model not start)
	//if !config.WeChat.Debug {
	//	OpenPlatform.StartRefreshAuthListTokenTicker()
	//}

	// open platform official account
	//OpenOfficialAccount = OpenPlatform.GetOfficialAccount(config.WeChat.OpenPlatformAppId)
	//
	//offConf := &offConfig.Config{
	//	AppID:          config.WeChat.OpenPlatformAppId,
	//	AppSecret:      config.WeChat.OpenPlatformAppSecret,
	//	Token:          config.WeChat.OpenPlatformToken,
	//	EncodingAESKey: config.WeChat.OpenPlatformAesKey,
	//	Cache:          redisCache,
	//}
	//MiniProgramAccount = wc.GetOfficialAccount(offConf)
	// start cron update ticket (local model not start)
	//if !config.WeChat.Debug {
	//	MiniProgramAccount.StartRefreshAccessTokenTicker()
	//}

	// mini program login
	miniConf := &miniConfig.Config{
		AppID:     config.WeChat.MiniProgramAppId,
		AppSecret: config.WeChat.MiniProgramAppSecret,
		Cache:     redisCache,
	}
	MiniProgram = wc.GetMiniProgram(miniConf)
}
