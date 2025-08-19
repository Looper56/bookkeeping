package service

import (
	"bookkeeping/internal/model"
	"bookkeeping/internal/repository"
	"context"
)

// LoginService login service
type LoginService struct {
	userService    *UserService
	userRepository *repository.UserRepository
	sessionService *SessionService
	mpService      *MiniProgramService
}

// NewLoginService init
func NewLoginService() *LoginService {
	return &LoginService{
		userService:    NewUserService(),
		userRepository: repository.NewUserRepository(),
		sessionService: NewSessionService(),
		mpService:      NewMiniProgramService(),
	}
}

// MPAuth miniProgram auth service
func (l *LoginService) MPAuth(ctx context.Context, code string) (*model.LoginState, error) {
	session, err := l.mpService.CodeToSession(code)
	if err != nil {
		return nil, err
	}
	defaultUser := &model.User{
		Channel:  model.WechatChannel,
		MPOpenID: session.OpenID,
		UnionID:  session.UnionID,
	}
	user, err := l.userService.UpsertUser(ctx, defaultUser)
	if err != nil {
		return nil, err
	}
	session.UserUID = user.UID
	session.IsWhite = user.IsWhite
	session.Mobile = user.Mobile
	token, err := l.sessionService.SaveSession(ctx, session)
	if err != nil {
		return nil, err
	}
	loginState := &model.LoginState{
		Token:     token.Token,
		OpenID:    user.MPOpenID,
		IsWhite:   user.IsWhite,
		HasMobile: user.Mobile != "",
	}
	return loginState, nil
}
