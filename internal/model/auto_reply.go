package model

// AutoReplyConfig 自动回复
type AutoReplyConfig struct {
	Id           int    `json:"id" gorm:"column:id"`
	Title        string `json:"title" gorm:"column:title"`
	Type         int    `json:"type" gorm:"column:type"`               // 1: keywords to reply 2: sub reply
	ReplyType    int    `json:"replay_type" gorm:"column:replay_type"` // reply type（1: text 2: img 3: img text）
	ImportKey    string `json:"import_key" gorm:"column:import_key"`   // keywords
	ReplyContent string `json:"replay" gorm:"column:replay"`           // reply content
	IsOnline     int    `json:"is_online" gorm:"column:is_online"`     // is active
	ChannelType  int    `json:"channel_type" gorm:"column:channel_type"`
}

// TableName 表名
func (AutoReplyConfig) TableName() string {
	return "auto_replay"
}
