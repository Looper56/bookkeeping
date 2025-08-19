package repository

import (
	"bookkeeping/internal/model"
	"bookkeeping/pkg/pagination"
	"context"
	"errors"
	"time"

	"github.com/Looper56/plugin/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserRepository ...
type UserRepository struct {
	mysql.Connector
}

// NewUserRepository init
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create ...
func (u *UserRepository) Create(ctx context.Context, tx *gorm.DB, user *model.User) error {
	return tx.Create(user).Error
}

// Update ...
func (u *UserRepository) Update(ctx context.Context, scope func(db *gorm.DB) *gorm.DB, user *model.User) error {
	return u.DB().WithContext(ctx).Model(&model.User{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid"}},
		UpdateAll: true,
	}).Scopes(scope).Create(&user).Error
}

type FindOneUserCondition struct {
	UID      string
	UnionID  string
	MPOpenID string
}

// FindOne find one by condition
func (u *UserRepository) FindOne(ctx context.Context, conn *FindOneUserCondition, fields ...string) (*model.User, error) {
	var user model.User
	sql := u.DB().WithContext(ctx).Where(conn).Where("is_cancel = ?", 0)
	if len(fields) > 0 {
		sql = sql.Select(fields)
	}
	sql.First(&user)
	if errors.Is(sql.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, sql.Error
}

// UserIsExist user exist or not
func (u *UserRepository) UserIsExist(ctx context.Context, scope func(db *gorm.DB) *gorm.DB) bool {
	var isExist bool
	u.DB().WithContext(ctx).Model(&model.User{}).Select("count(*) = 0").Scopes(scope).Find(&isExist)
	return isExist
}

// UserUnSubscribe user unsubscribe
func (u *UserRepository) UserUnSubscribe(ctx context.Context, now *time.Time, openID string) error {
	return u.DB().
		WithContext(ctx).
		Model(&model.User{}).
		Where(&model.User{OfficialAccountID: openID}).
		UpdateColumns(map[string]interface{}{
			"is_sub_oa":        0,
			"unfollow_oa_time": now,
		}).Error
}

// UserRolePagination ...
type UserRolePagination *pagination.LengthAwarePagination[*model.User]

// RolePermissionList ...
//func (u *UserRepository) RolePermissionList(ctx context.Context, params *request.RolePermissionRequest) (
//	UserRolePagination, error) {
//	var rolePermission []*model.RolePermission
//	pag := pagination.NewLengthAwarePagination[*model.RolePermission](
//		params.GetPageSize(), params.GetPage(), "created_at desc")
//	sql := u.DB().Where("type in (?)", []int8{model.OAUser, model.EmailUser})
//	sql = sql.WithContext(ctx).Scopes(pagination.LengthAwarePaginate(rolePermission, pag, sql))
//	if params.HasAccountKeyWords() {
//		sql = sql.Scopes(UserAccountKeywordScope(params.GetAccountKeyWords()))
//	}
//	if params.HasUserNameKeyWords() {
//		sql = sql.Scopes(UserNameKeywordScope(params.GetUserNameKeyWords()))
//	}
//	sql = sql.Find(&rolePermission)
//	pag.SetRows(rolePermission)
//	return pag, errors.Wrapf(sql.Error, "失败: %+v", params)
//}
