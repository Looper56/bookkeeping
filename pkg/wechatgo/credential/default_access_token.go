package credential

import (
	"bookkeeping/pkg/wechatgo/cache"
	"bookkeeping/pkg/wechatgo/util"
	"fmt"
)

const (
	// CacheKeyOfficialAccountPrefix cache key前缀
	CacheKeyOfficialAccountPrefix = "wechat_officialaccount_"
)

// DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	appID     string
	appSecret string
	cache     cache.Cache
}

// NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(appID, appSecret string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is need")
	}
	return &DefaultAccessToken{
		appID:     appID,
		appSecret: appSecret,
		cache:     cache,
	}
}

// ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetAccessToken 获取access_token,仅从cache中获取,token是通过定时器更新的
func (ak *DefaultAccessToken) GetAccessToken() (accessToken string, err error) {
	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", CacheKeyOfficialAccountPrefix, ak.appID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
	}
	return
}
