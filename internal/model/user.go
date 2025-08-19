package model

import (
	"time"
)

var (
	WechatChannel int8 = 1
)

var (
	IsWhite int8 = 1
)

var (
	SubOfficialAccount = 1
)

// User ...
type User struct {
	// ID
	ID  int64  `json:"id,omitempty" gorm:"column:id"`
	UID string `json:"uid,omitempty" gorm:"column:uid"`
	// officialAccount open id
	OfficialAccountID string `json:"official_account_id,omitempty" gorm:"column:official_account_id"`
	// miniProgram open id
	MPOpenID string `json:"mp_open_id,omitempty" gorm:"column:mp_open_id"`
	// the same id for miniProgram and officialAccount
	UnionID string `json:"union_id,omitempty" gorm:"column:union_id"`
	Channel int8   `json:"channel,omitempty" gorm:"column:channel"`

	// profile
	NickName  string `json:"nick_name" gorm:"column:nick_name"`
	AvatarURL string `json:"avatar_url" gorm:"column:avatar_url"`
	Gender    int32  `json:"gender" gorm:"column:gender"`
	Mobile    string `json:"mobile" gorm:"column:mobile"`
	Province  string `json:"province" gorm:"column:province"`
	Country   string `json:"country" gorm:"column:country"`
	City      string `json:"city" gorm:"column:city"`

	// status control
	IsWhite              int8 `json:"is_white,omitempty" gorm:"column:is_white"`
	IsCanceled           int  `json:"is_cancel,omitempty" gorm:"column:is_cancel"` // is canceled account
	IsSubOfficialAccount int  `json:"is_sub_oa,omitempty" gorm:"column:is_sub_oa"` // is subscribed official account

	// time
	LastActiveTime  *time.Time `json:"last_active_time,omitempty" gorm:"column:last_active_time"`
	SubscribeTime   *time.Time `json:"sub_oa_time,omitempty" gorm:"column:sub_oa_time"` // first time sub
	UnSubscribeTime *time.Time `json:"unfollow_oa_time,omitempty" gorm:"column:unfollow_oa_time"`
}

// TableName ...
func (u User) TableName() string {
	return "user"
}
