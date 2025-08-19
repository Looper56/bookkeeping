package wechatgo

import (
	"bookkeeping/pkg/wechatgo/cache"
	"bookkeeping/pkg/wechatgo/miniprogram"
	miniConfig "bookkeeping/pkg/wechatgo/miniprogram/config"
	"bookkeeping/pkg/wechatgo/officialaccount"
	offConfig "bookkeeping/pkg/wechatgo/officialaccount/config"
	"bookkeeping/pkg/wechatgo/openplatform"
	openConfig "bookkeeping/pkg/wechatgo/openplatform/config"
	"bookkeeping/pkg/wechatgo/pay"
	payConfig "bookkeeping/pkg/wechatgo/pay/config"
)

// WeChat struct
type WeChat struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *WeChat {
	return &WeChat{}
}

// SetCache 设置cache
func (wc *WeChat) SetCache(cache cache.Cache) {
	wc.cache = cache
}

// GetOfficialAccount 获取微信公众号实例
func (wc *WeChat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *WeChat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *WeChat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *WeChat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	return openplatform.NewOpenPlatform(cfg)
}
