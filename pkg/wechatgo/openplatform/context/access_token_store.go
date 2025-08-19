package context

import (
	"bookkeeping/pkg/wechatgo/api"
	"bookkeeping/pkg/wechatgo/util"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"

	log "github.com/sirupsen/logrus"
)

// saveVerifyTicket 存储verifyTicket
func (ctx *Context) saveVerifyTicket(verifyTicket string) error {
	now := time.Now()
	var token ComponentAppTokens
	res := ctx.DB.Model(&ComponentAppTokens{}).
		Select([]string{"id"}).
		Where(&ComponentAppTokens{Appid: ctx.AppID}).
		First(&token)
	updateInfo := ComponentAppTokens{
		Appid:                 ctx.AppID,
		Appsecret:             ctx.AppSecret,
		ComponentVerifyTicket: verifyTicket,
		TicketUpdatedAt:       &now,
		Status:                1,
		UpdatedAt:             &now,
	}
	// use gorm.io/gorm
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		updateInfo.InfoType = "component_verify_ticket"
		updateInfo.CreatedAt = &now
		res = ctx.DB.Create(&updateInfo)
		if res.Error != nil {
			return res.Error
		}
	} else {
		res = ctx.DB.Model(&token).Updates(&updateInfo)
		if res.Error != nil {
			return res.Error
		}
	}

	// use jinzhu/gorm
	//if res.RecordNotFound() {
	//	updateInfo.InfoType = "component_verify_ticket"
	//	updateInfo.CreatedAt = &now
	//	res := ctx.DB.Create(&updateInfo)
	//	if res.Error != nil {
	//		return res.Error
	//	}
	//} else {
	//	res := ctx.DB.Model(&token).Update(&updateInfo)
	//	if res.Error != nil {
	//		return res.Error
	//	}
	//}

	cacheKey := fmt.Sprintf("verify_ticket_%s", ctx.AppID)
	err := ctx.Cache.Set(cacheKey, verifyTicket, 12*time.Hour)
	if err != nil {
		return err
	}

	return nil
}

// getComponentAccessTokenNotExpired 从缓存中获取component access token
func (ctx *Context) getComponentAccessTokenNotExpired() *ComponentAccessToken {
	accessTokenCacheKey := fmt.Sprintf("component_access_token_%s", ctx.AppID)
	at := ctx.Cache.Get(accessTokenCacheKey)
	if at == nil {
		return nil
	}

	atStore := &componentAccessTokenStore{}
	atStr := at.(string)
	err := json.Unmarshal([]byte(atStr), atStore)
	// 1个半小时更新一次
	if err == nil && time.Now().Sub(atStore.ExpiresAt) <= 90*time.Minute {
		return atStore.Token
	}
	return nil
}

// updateComponentAccessToken 更新component access token
func (ctx *Context) updateComponentAccessToken(verifyTicket string) (*ComponentAccessToken, error) {
	accessTokenCacheKey := fmt.Sprintf("component_access_token_%s", ctx.AppID)
	now := time.Now()

	body := map[string]string{
		"component_appid":         ctx.AppID,
		"component_appsecret":     ctx.AppSecret,
		"component_verify_ticket": verifyTicket,
	}
	respBody, err := util.PostJSON(api.ComponentAccessTokenUri(), body)
	if err != nil {
		return nil, err
	}

	at := &ComponentAccessToken{}
	if err := json.Unmarshal(respBody, at); err != nil {
		return nil, err
	}
	atBytes, err := json.Marshal(componentAccessTokenStore{
		Token:     at,
		ExpiresAt: now,
	})
	if err != nil {
		return nil, err
	}
	if at.AccessToken == "" {
		err = fmt.Errorf("get component access token fail: %s", string(respBody))
		log.Error(err.Error())
		return nil, err
	}

	nextUpdateTime := now.Add(time.Duration(at.ExpiresIn-1500) * time.Second)
	res := ctx.DB.Model(&ComponentAppTokens{}).
		Where(&ComponentAppTokens{Appid: ctx.AppID}).
		Updates(&ComponentAppTokens{
			ComponentAccessToken: at.AccessToken,
			ExpiresIn:            at.ExpiresIn,
			TokenUpdatedAt:       &now,
			NextUpdateTime:       &nextUpdateTime,
			UpdatedAt:            &now,
		})
	//res := ctx.DB.Model(&ComponentAppTokens{}).
	//	Where(&ComponentAppTokens{Appid: ctx.AppID}).
	//	Update(&ComponentAppTokens{
	//		ComponentAccessToken: at.AccessToken,
	//		ExpiresIn:            at.ExpiresIn,
	//		TokenUpdatedAt:       &now,
	//		NextUpdateTime:       &nextUpdateTime,
	//		UpdatedAt:            &now,
	//	})
	if res.Error != nil {
		return nil, res.Error
	}

	err = ctx.Cache.Set(accessTokenCacheKey, string(atBytes), time.Duration(at.ExpiresIn)*time.Second)
	if err != nil {
		return nil, nil
	}

	return at, nil
}

// SaveAuthInfo 存储授权码获取的授权信息
func (ctx *Context) SaveAuthInfo(authInfo *AuthBaseInfo) error {
	now := time.Now()
	var authorizerToken AuthorizerAppTokens
	res := ctx.DB.Model(&AuthorizerAppTokens{}).Where(&AuthorizerAppTokens{
		ComponentAppid:  ctx.AppID,
		AuthorizerAppid: authInfo.Appid,
	}).First(&authorizerToken)
	nextUpdateTime := now.Add(time.Duration(authInfo.ExpiresIn-1500) * time.Second)
	updateInfo := AuthorizerAppTokens{
		AuthorizerAccessToken:  authInfo.AccessToken,
		ExpiresIn:              authInfo.ExpiresIn,
		AuthorizerRefreshToken: authInfo.RefreshToken,
		TokenUpdatedAt:         &now,
		NextUpdateTime:         &nextUpdateTime,
		Status:                 1,
		UpdatedAt:              &now,
	}
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		updateInfo.ComponentAppid = ctx.AppID
		updateInfo.AuthorizerAppid = authInfo.Appid
		res = ctx.DB.Create(&updateInfo)
		if res.Error != nil {
			return res.Error
		}
	} else {
		res = ctx.DB.Model(&authorizerToken).Updates(&updateInfo)
		if res.Error != nil {
			return res.Error
		}
	}
	//if res.RecordNotFound() {
	//	updateInfo.ComponentAppid = ctx.AppID
	//	updateInfo.AuthorizerAppid = authInfo.Appid
	//	res = ctx.DB.Create(&updateInfo)
	//	if res.Error != nil {
	//		return res.Error
	//	}
	//} else {
	//	res = ctx.DB.Model(&authorizerToken).Update(&updateInfo)
	//	if res.Error != nil {
	//		return res.Error
	//	}
	//}

	err := ctx.CacheAuthrAccessToken(authInfo.Appid, authInfo.AccessToken, authInfo.ExpiresIn)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAuthrAccessToken 更新授权方AccessToken
func (ctx *Context) UpdateAuthrAccessToken(authrAccessToken *AuthrAccessToken) error {
	now := time.Now()
	nextUpdateTime := now.Add(time.Duration(authrAccessToken.ExpiresIn-1500) * time.Second)
	res := ctx.DB.Model(&AuthorizerAppTokens{}).Where(&AuthorizerAppTokens{
		ComponentAppid:  ctx.AppID,
		AuthorizerAppid: authrAccessToken.Appid,
	}).Updates(AuthorizerAppTokens{
		AuthorizerAccessToken:  authrAccessToken.AccessToken,
		ExpiresIn:              authrAccessToken.ExpiresIn,
		AuthorizerRefreshToken: authrAccessToken.RefreshToken,
		TokenUpdatedAt:         &now,
		NextUpdateTime:         &nextUpdateTime,
		UpdatedAt:              &now,
	})
	//res := ctx.DB.Model(&AuthorizerAppTokens{}).Where(&AuthorizerAppTokens{
	//	ComponentAppid:  ctx.AppID,
	//	AuthorizerAppid: authrAccessToken.Appid,
	//}).Update(AuthorizerAppTokens{
	//	AuthorizerAccessToken:  authrAccessToken.AccessToken,
	//	ExpiresIn:              authrAccessToken.ExpiresIn,
	//	AuthorizerRefreshToken: authrAccessToken.RefreshToken,
	//	TokenUpdatedAt:         &now,
	//	NextUpdateTime:         &nextUpdateTime,
	//	UpdatedAt:              &now,
	//})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// CacheAuthrAccessToken 缓存授权方AccessToken
func (ctx *Context) CacheAuthrAccessToken(appID, accessToken string, expiresIn int64) error {
	authrTokenKey := "authorizer_access_token_" + appID
	if err := ctx.Cache.Set(authrTokenKey, accessToken, time.Duration(expiresIn)*time.Second); err != nil {
		return err
	}
	return nil
}

// StartRefreshAuthListTokenTicker 启动批量获取（刷新）授权公众号或小程序的接口调用凭据（令牌）定时器
func (ctx *Context) StartRefreshAuthListTokenTicker() {
	go func() {
		// 先更新一次
		ctx.StartRefreshAuthrListToken()

		t := time.NewTicker(1 * time.Minute)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				ctx.StartRefreshAuthrListToken()
			}
		}
	}()
}

// StartRefreshAuthrListToken 批量获取（刷新）授权公众号或小程序的接口调用凭据（令牌）
func (ctx *Context) StartRefreshAuthrListToken() {
	go func() {
		lockKey := fmt.Sprintf("refersh_token_%s", ctx.AppID)
		// 分布式锁
		ok, _ := ctx.Cache.Lock(lockKey, 10*time.Second)
		defer ctx.Cache.UnLock(lockKey)
		if ok {
			err := ctx.RefreshAuthrListToken(ctx.AppID)
			if err != nil {
				log.Errorf("refresh authorizer list token fail.")
			}
		}
		log.Info("finish refresh component access token")
	}()
}

// RefreshAuthrListToken 批量获取（刷新）授权公众号或小程序的接口调用凭据（令牌）
func (ctx *Context) RefreshAuthrListToken(appID string) error {
	var authorizerTokens []AuthorizerAppTokens
	res := ctx.DB.Model(&AuthorizerAppTokens{}).
		Select([]string{"authorizer_appid", "authorizer_refresh_token"}).
		Where("component_appid = ? AND next_update_time <= ? AND status = 1", appID, time.Now()).
		Find(&authorizerTokens)
	if res.Error != nil {
		return res.Error
	}
	for _, authorizerToken := range authorizerTokens {
		_, err := ctx.RefreshAuthrToken(authorizerToken.AuthorizerAppid, authorizerToken.AuthorizerRefreshToken)
		if err != nil {
			log.Errorf("authorizer access token update fail, appid: %s, err: %s",
				authorizerToken.AuthorizerAppid, err.Error())
		} else {
			log.Infof("authorizer access token update success, appid: %s", authorizerToken.AuthorizerAppid)
		}
	}
	return nil
}
