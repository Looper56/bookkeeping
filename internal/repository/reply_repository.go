package repository

import (
	"bookkeeping/internal/model"
	"context"

	"github.com/Looper56/plugin/mysql"
)

// ReplyRepository reply resp
type ReplyRepository struct {
	mysql.Connector
}

// NewReplyRepository ...
func NewReplyRepository() *ReplyRepository {
	return &ReplyRepository{}
}

// FindOneQRCodeCondition find condition
type FindOneQRCodeCondition struct {
	SceneStr string
}

// QRCodeDetail ...
func (r *ReplyRepository) QRCodeDetail(ctx context.Context, conn *FindOneQRCodeCondition, field ...string) (
	*model.QRCodeConfig, error) {
	var qrCode model.QRCodeConfig
	sql := r.DB().WithContext(ctx).Select(field).Where(conn).First(&qrCode)
	return &qrCode, sql.Error
}

// FindAutoReplyCondition find condition
type FindAutoReplyCondition struct {
	Type        int
	ChannelType int
}

// AutoReplyDetail ...
func (r *ReplyRepository) AutoReplyDetail(ctx context.Context, conn *FindAutoReplyCondition) (
	*model.AutoReplyConfig, error) {
	var autoReply model.AutoReplyConfig
	sql := r.DB().WithContext(ctx).Where(conn).First(&autoReply)
	return &autoReply, sql.Error
}
