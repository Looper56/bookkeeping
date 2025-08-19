package service

import (
	"bookkeeping/internal/model"
	"bookkeeping/internal/repository"
	"bookkeeping/pkg/wechatgo/officialaccount/message"
	"context"
	"encoding/json"
	"strings"
)

type ReplyService struct {
	customManger    *message.Manager
	replyRepository *repository.ReplyRepository
}

func NewReplyService() *ReplyService {
	return &ReplyService{
		replyRepository: repository.NewReplyRepository(),
	}
}

var QRCodeFields = []string{"id", "replay_type", "replay_content"}

// GetQRCodeByEventKey get QR code setting
func (r *ReplyService) GetQRCodeByEventKey(ctx context.Context, eventKey string) (*model.QRCodeConfig, error) {
	keys := strings.Split(eventKey, "_")
	sceneStr := keys[1]
	qrCode, err := r.replyRepository.QRCodeDetail(ctx,
		&repository.FindOneQRCodeCondition{SceneStr: sceneStr},
		QRCodeFields...)
	if err != nil {
		return nil, err
	}
	return qrCode, nil
}

// TextAutoReply text
func (r *ReplyService) TextAutoReply(ctx context.Context, toUser string, eventType int) (*message.Reply, error) {
	autoReply, err := r.replyRepository.AutoReplyDetail(ctx,
		&repository.FindAutoReplyCondition{Type: eventType, ChannelType: 1},
	)
	if err != nil {
		return nil, err
	}
	return r.ReplyContent(toUser, autoReply.ReplyContent, autoReply.ReplyType)
}

var MsgTypeImgText message.MsgType = "mpnews"

// ReplyContent unify reply
func (r *ReplyService) ReplyContent(toUser, content string, replyType int) (*message.Reply, error) {
	switch replyType {
	// text
	case 1:
		return &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: message.NewText(content),
		}, nil
	// img
	case 2:
		return &message.Reply{
			MsgType: message.MsgTypeImage,
			MsgData: message.NewImage(content),
		}, nil
	case 3:
		// img text
		news := &message.CustomerMessage{
			ToUser:  toUser,
			Msgtype: MsgTypeImgText,
			Mpnews:  &message.MediaResource{MediaID: content},
		}
		err := r.customManger.Send(news)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case 4:
		// customer msg mini program
		var msg message.MediaMiniprogrampage
		err := json.Unmarshal([]byte(content), &msg)
		if err != nil {
			return nil, err
		}
		miniMsg := &message.CustomerMessage{
			ToUser:          toUser,
			Msgtype:         message.MsgTypeMiniprogrampage,
			Miniprogrampage: &msg,
		}
		err = r.customManger.Send(miniMsg)
		if err != nil {
			return nil, err
		}
		return nil, nil
	default:
		return nil, nil
	}
}
