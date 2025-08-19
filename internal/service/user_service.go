package service

import (
	"bookkeeping/internal/app/http/request"
	"bookkeeping/internal/model"
	"bookkeeping/internal/repository"
	"bookkeeping/internal/scope"
	"bookkeeping/pkg/util"
	"bookkeeping/pkg/wechatgo/openplatform/officialaccount"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Looper56/plugin/mysql"
)

// UserService user service
type UserService struct {
	mysql.Connector
	userRepository  *repository.UserRepository
	mpService       *MiniProgramService
	officialAccount *officialaccount.OfficialAccount
}

// NewUserService init
func NewUserService() *UserService {
	var openOfficeAccount *officialaccount.OfficialAccount
	return &UserService{
		userRepository:  repository.NewUserRepository(),
		mpService:       NewMiniProgramService(),
		officialAccount: openOfficeAccount,
	}
}

// UpsertUser update or create user in auth
func (u *UserService) UpsertUser(ctx context.Context, latestUser *model.User) (*model.User, error) {
	userExist, err := u.userRepository.FindOne(ctx, &repository.FindOneUserCondition{UnionID: latestUser.UnionID})
	if err != nil {
		return nil, err
	}
	tx := u.DB().WithContext(ctx).Begin()
	err = tx.Error
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	now := time.Now()
	if userExist == nil {
		latestUser.UID = util.UniqueID()
		latestUser.AvatarURL = "https://dwk-1253502009.file.myqcloud.com/cdn/img/default-avatar.jpg"
		latestUser.NickName = fmt.Sprintf("%s-%s", "user", util.RandStr(5))
		latestUser.LastActiveTime = &now
		latestUser.IsWhite = model.IsWhite
		err = u.userRepository.Create(ctx, tx, latestUser)
		if err != nil {
			return nil, err
		}
		return latestUser, nil
	} else {
		latestUser.LastActiveTime = &now
		condition := []*scope.UpdateScope{
			{
				Condition: "mp_open_id",
				Value:     latestUser.MPOpenID,
			},
		}
		err = u.userRepository.Update(ctx, scope.Update(condition...), latestUser)
		if err != nil {
			return nil, err
		}
		return userExist, nil
	}
}

var ErrSessionKeyInvalid = errors.New("invalid padding on input")

// MPUserInfo miniProgram user info
func (u *UserService) MPUserInfo(ctx context.Context, session *model.Session, encryptedData, iv string) error {
	user, err := u.mpService.EncryptUserInfo(session.SessionKey, encryptedData, iv)
	if err != nil {
		if errors.Is(err, ErrSessionKeyInvalid) {
			return ErrSessionKeyInvalid
		}
		return err
	}

	condition := []*scope.UpdateScope{
		{
			Condition: "union_id",
			Value:     session.UnionID,
		},
	}
	err = u.userRepository.Update(ctx, scope.Update(condition...), user)
	if err != nil {
		return err
	}
	return nil
}

// GetOfficialAccountInfo get official account info
func (u *UserService) GetOfficialAccountInfo(openID string) (*model.User, error) {
	user, err := u.officialAccount.GetUser().GetUserInfo(openID)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Channel:           model.WechatChannel,
		OfficialAccountID: user.OpenID,
		UnionID:           user.UnionID,
		NickName:          user.Nickname,
		AvatarURL:         user.Headimgurl,
		Gender:            user.Sex,
		Country:           user.Country,
		Province:          user.Province,
	}, nil
}

func (u *UserService) UserInfo(ctx context.Context, session *model.Session) (*model.User, error) {
	return u.userRepository.FindOne(ctx, &repository.FindOneUserCondition{MPOpenID: session.OpenID})
}

func (u *UserService) UpdateUser(ctx context.Context, mpOpenID string, params request.UpdateUserRequest) error {
	user, err := u.userRepository.FindOne(ctx, &repository.FindOneUserCondition{MPOpenID: mpOpenID})
	if err != nil {
		return err
	}
	user.NickName = params.NickName
	user.Gender = params.Gender
	user.Mobile = params.Mobile
	user.Country = params.Country
	user.Province = params.Province
	user.City = params.City
	fmt.Printf("new user modle: %+v\n", user)

	condition := []*scope.UpdateScope{
		{
			Condition: "mp_open_id",
			Value:     mpOpenID,
		},
	}
	return u.userRepository.Update(ctx, scope.Update(condition...), user)
}
