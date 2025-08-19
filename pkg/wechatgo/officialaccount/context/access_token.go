package context

import (
	"bookkeeping/pkg/wechatgo/api"
	"bookkeeping/pkg/wechatgo/credential"
	"bookkeeping/pkg/wechatgo/util"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// StartRefreshAccessTokenTicker 启动定时更新access_token，由调用方触发
func (ctx *Context) StartRefreshAccessTokenTicker() {
	go func() {
		// 先更新一次
		ctx.StartRefreshAccessToken()

		t := time.NewTicker(1 * time.Minute)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				ctx.StartRefreshAccessToken()
			}
		}
	}()
}

// StartRefreshAccessToken 获取（刷新）授权公众号或小程序的接口调用凭据（令牌）(异步）
func (ctx *Context) StartRefreshAccessToken() {
	go func() {
		lockKey := fmt.Sprintf("official_refersh_token_%s", ctx.AppID)
		// 分布式锁
		ok, err := ctx.Cache.Lock(lockKey, 10*time.Second)
		if err != nil {
			log.Infof("finish refresh official account %s access token failed: connected redis failed %v", ctx.AppID, err)
			return
		}
		defer ctx.Cache.UnLock(lockKey)
		if !ok {
			log.Infof("skip refresh official account %s access token: locked", ctx.AppID)
			return
		}
		err = ctx.RefreshAccessToken()
		if err != nil {
			log.Errorf("refresh official account %s access token fail. err: %s", ctx.AppID, err.Error())
		}
		log.Infof("finish refresh official account %s access token", ctx.AppID)
	}()
}

// RefreshAccessToken 获取（刷新）授权公众号或小程序的接口调用凭据（令牌）（同步）
func (ctx *Context) RefreshAccessToken() error {
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s",
		credential.CacheKeyOfficialAccountPrefix, ctx.AppID)
	// ttl key用来判断什么时候更新token
	accessTokenTTLKey := fmt.Sprintf("%s_access_token_ttl_%s",
		credential.CacheKeyOfficialAccountPrefix, ctx.AppID)

	val := ctx.Cache.Get(accessTokenCacheKey)
	notExpired := ctx.Cache.IsExist(accessTokenTTLKey)
	// 还未过期
	if val != nil && notExpired {
		log.Infof("skip refresh official account %s access token: not expired, token: %v", ctx.AppID, val)
		return nil
	}

	// cache失效，从微信服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err := ctx.GetTokenFromServer(ctx.AppID, ctx.AppSecret)
	if err != nil {
		return err
	}
	log.Infof("get official account %s access token(%v), expire(%v)",
		ctx.AppID, resAccessToken.AccessToken, resAccessToken.ExpiresIn)

	err = ctx.Cache.Set(accessTokenCacheKey, resAccessToken.AccessToken,
		time.Duration(resAccessToken.ExpiresIn)*time.Second)
	if err != nil {
		return err
	}
	expires := resAccessToken.ExpiresIn - 1500
	err = ctx.Cache.Set(accessTokenTTLKey, "1", time.Duration(expires)*time.Second)
	if err != nil {
		return err
	}

	log.Infof("refresh official account %s access token from server.", ctx.AppID)
	return nil
}

// GetTokenFromServer 强制从微信服务器获取token
func (ctx *Context) GetTokenFromServer(appID, appSecret string) (resAccessToken ResAccessToken, err error) {
	url := api.AccessTokenUri(appID, appSecret)
	var body []byte
	body, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.ErrMsg != "" {
		err = fmt.Errorf("get access_token error: errcode=%v, errormsg=%v",
			resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}
