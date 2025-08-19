package context

import "time"

// ComponentAppTokens 第三方开放平台账号
type ComponentAppTokens struct {
	// ID 物理主键
	ID uint32 `json:"id" gorm:"primaryKey"`
	// Appid 第三方平台appid
	Appid string `json:"appid"`
	// Appsecret 第三方平台appsecret
	Appsecret string `json:"appsecret"`
	// InfoType 信息类型，现在只有component_verify_ticket
	InfoType string `json:"info_type"`
	// ComponentVerifyTicket ticket内容，在第三方平台创建审核通过后，微信服务器会向其“授权事件接收URL”
	// 每隔10分钟定时推送component_verify_ticket。
	ComponentVerifyTicket string `json:"component_verify_ticket"`
	// TicketUpdatedAt component_verify_ticket更新时间
	TicketUpdatedAt *time.Time `json:"ticket_updated_at"`
	// ComponentAccessToken 第三方平台component_access_token是第三方平台的下文中接口的调用凭据，
	// 也叫做令牌（component_access_token）。每个令牌是存在有效期（2小时）的，且令牌的调用不是无限制的，
	// 请第三方平台做好令牌的管理，在令牌快过期时（比如1小时50分）再进行刷新
	ComponentAccessToken string `json:"component_access_token"`
	// ExpiresIn 有效期
	ExpiresIn int64 `json:"expires_in"`
	// TokenUpdatedAt component_access_token更新时间
	TokenUpdatedAt *time.Time `json:"token_updated_at"`
	// NextUpdateTime component_access_token下次更新时间
	NextUpdateTime *time.Time `json:"next_update_time"`
	// Status 有效情况，1生效，0失效，默认1
	Status uint8 `json:"status"`
	// CreatedAt 创建时间
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt *time.Time `json:"updated_at"`
	// Remark 备注信息
	Remark string `json:"remark"`
}

// TableName 表名
func (ComponentAppTokens) TableName() string {
	return "component_app_tokens"
}

// AuthorizerAppTokens 授权账号
type AuthorizerAppTokens struct {
	// ID 物理主键
	ID uint32 `json:"id" gorm:"primaryKey"`
	// ComponentAppid 第三方平台appid
	ComponentAppid string `json:"component_appid"`
	// AuthorizerOriginID 授权方原始ID
	AuthorizerOriginID string `json:"authorizer_origin_id"`
	// AuthorizerAppid 授权方appid
	AuthorizerAppid string `json:"authorizer_appid"`
	// AuthorizerAccessToken 授权方接口调用凭据（在授权的公众号或小程序具备API权限时，才有此返回值），也简称为令牌
	AuthorizerAccessToken string `json:"authorizer_access_token"`
	// ExpiresIn 有效期（在授权的公众号或小程序具备API权限时，才有此返回值）
	ExpiresIn int64 `json:"expires_in"`
	// AuthorizerRefreshToken 接口调用凭据刷新令牌（在授权的公众号具备API权限时，才有此返回值），
	// 刷新令牌主要用于第三方平台获取和刷新已授权用户的access_token，只会在授权时刻提供，请妥善保存。
	// 一旦丢失，只能让用户重新授权，才能再次拿到新的刷新令牌
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	// FuncInfo 授权给开发者的权限集列表
	FuncInfo string `json:"func_info"`
	// AuthorizerInfo 授权方的帐号基本信息
	AuthorizerInfo string `json:"authorizer_info"`
	// TokenUpdatedAt authorizer_access_token更新时间
	TokenUpdatedAt *time.Time `json:"token_updated_at"`
	// NextUpdateTime authorizer_access_token下次更新时间
	NextUpdateTime *time.Time `json:"next_update_time"`
	// Status 有效情况，1生效，0失效，默认1
	Status uint8 `json:"status"`
	// UpdatedAt 更新时间
	UpdatedAt *time.Time `json:"updated_at"`
	// CreatedAt 创建时间
	CreatedAt *time.Time `json:"created_at"`
}

// TableName 表名
func (AuthorizerAppTokens) TableName() string {
	return "authorizer_app_tokens"
}
