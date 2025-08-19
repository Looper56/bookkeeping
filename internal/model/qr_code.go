package model

// QRCodeConfig 小程序二维码
type QRCodeConfig struct {
	Id           int    `json:"id" gorm:"column:id"`
	SceneStr     string `json:"scene_str" gorm:"column:scene_str"`           // user set id to generate QR code
	ReplyType    int    `json:"replay_type" gorm:"column:replay_type"`       // auto reply type（1: article 2: img 3: img text）
	ReplyContent string `json:"replay_content" gorm:"column:replay_content"` // content to auto reply
}

// TableName 表名
func (QRCodeConfig) TableName() string {
	return "qrcode"
}
