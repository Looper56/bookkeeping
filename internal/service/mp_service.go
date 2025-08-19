package service

import (
	"bookkeeping/config"
	"bookkeeping/internal/model"
	"bookkeeping/pkg/wechatgo/miniprogram"
)

// MiniProgramService miniProgram service with API
type MiniProgramService struct {
	miniProgram *miniprogram.MiniProgram
}

// NewMiniProgramService init
func NewMiniProgramService() *MiniProgramService {
	//var MP *miniprogram.MiniProgram
	return &MiniProgramService{
		miniProgram: config.MiniProgram,
	}
}

// CodeToSession get session by auth code
func (m *MiniProgramService) CodeToSession(code string) (*model.Session, error) {
	auth := m.miniProgram.GetAuth()
	sessionRes, err := auth.Code2Session(code)
	if err != nil {
		return nil, err
	}
	session := &model.Session{
		Channel:    model.WechatChannel,
		IsWhite:    0,
		OpenID:     sessionRes.OpenID,
		UnionID:    sessionRes.UnionID,
		SessionKey: sessionRes.SessionKey,
	}
	return session, nil
}

// EncryptUserInfo encrypted user info from session
func (m *MiniProgramService) EncryptUserInfo(sessionKey, encryptedData, iv string) (*model.User, error) {
	user, err := m.miniProgram.GetEncryptor().Decrypt(sessionKey, encryptedData, iv)
	//logger.Info(fmt.Sprintf("encryptor.AppID: %+v", m.miniProgram.GetEncryptor().AppID))
	//logger.Info(fmt.Sprintf("sessionKey %+v", sessionKey))
	//logger.Info(fmt.Sprintf("encryptedData %+v", encryptedData))
	//logger.Info(fmt.Sprintf("iv %+v", iv))
	if err != nil {
		return nil, err
	}
	return &model.User{
		Channel:   model.WechatChannel,
		MPOpenID:  user.OpenID,
		UnionID:   user.UnionID,
		AvatarURL: user.AvatarURL,
		NickName:  user.NickName,
		City:      user.City,
	}, nil
}
