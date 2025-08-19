package service

import (
	"bookkeeping/internal/model"
	"bookkeeping/internal/repository"
	"bookkeeping/pkg/wechatgo/officialaccount/message"
	"bookkeeping/pkg/wechatgo/officialaccount/server"
	"context"
	"fmt"
	"time"

	"github.com/Looper56/plugin/logger"
)

// OfficialAccountService ...
type OfficialAccountService struct {
	server         *server.Server
	userService    *UserService
	userRepository *repository.UserRepository
	replyService   *ReplyService
	eventHandler   map[message.EventType]func(msg message.MixMessage) *message.Reply
	msgHandler     map[message.MsgType]func(msg message.MixMessage) *message.Reply
}

// NewOfficialAccountService init officialAccount service
func NewOfficialAccountService(server *server.Server) *OfficialAccountService {
	return &OfficialAccountService{
		server:         server,
		userService:    NewUserService(),
		userRepository: repository.NewUserRepository(),
		replyService:   NewReplyService(),
		eventHandler:   make(map[message.EventType]func(msg message.MixMessage) *message.Reply),
		msgHandler:     make(map[message.MsgType]func(msg message.MixMessage) *message.Reply),
	}
}

// OnEvent  msg event handle
func (o *OfficialAccountService) OnEvent(eventType message.EventType, cb func(msg message.MixMessage) *message.Reply) {
	o.eventHandler[eventType] = cb
}

// OnMsg msg event handle
func (o *OfficialAccountService) OnMsg(msgType message.MsgType, cb func(msg message.MixMessage) *message.Reply) {
	o.msgHandler[msgType] = cb
}

// Serve unify msg event handle
func (o *OfficialAccountService) Serve() error {
	var err error
	// only handle verify_ticket event
	o.server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		var reply *message.Reply
		// event handle
		if msg.MsgType == message.MsgTypeEvent {
			eventHandler, exist := o.eventHandler[msg.Event]
			if exist {
				reply = eventHandler(msg)
			}
		} else {
			msgHandler, exist := o.msgHandler[msg.MsgType]
			if exist {
				reply = msgHandler(msg)
			}
		}
		return reply
	})
	err = o.server.Serve()
	if err != nil {
		return err
	}

	err = o.server.Send()
	if err != nil {
		return err
	}
	return nil
}

// UserSubscribe user subscribe
func (o *OfficialAccountService) UserSubscribe(msg message.MixMessage) *message.Reply {
	var err error
	var reply *message.Reply
	ctx := context.Background()
	fromOpenID := string(msg.FromUserName)

	// user register
	user, err := o.userService.GetOfficialAccountInfo(fromOpenID)
	if err != nil {
		logger.Error(fmt.Sprintf("official account login err: %+v", err))
	}

	now := time.Now()
	user.IsSubOfficialAccount = model.SubOfficialAccount
	user.SubscribeTime = &now
	_, err = o.userService.UpsertUser(ctx, user)
	if err != nil {
		logger.Error(fmt.Sprintf("official account login err: %+v", err))
	}

	if msg.EventKey != "" {
		// scan QR code with param
		qrCode, qrCodeErr := o.replyService.GetQRCodeByEventKey(ctx, msg.EventKey)
		if qrCodeErr != nil {
			return nil
		}

		reply, err = o.replyService.ReplyContent(fromOpenID, qrCode.ReplyContent, qrCode.ReplyType)
	} else {
		reply, err = o.replyService.TextAutoReply(ctx, fromOpenID, 2)
	}

	if err != nil {
		logger.Error(fmt.Sprintf("official account reply err: %+v", err))
	}
	return reply
}

// UserUnSubscribe ...
func (o *OfficialAccountService) UserUnSubscribe(msg message.MixMessage) *message.Reply {
	ctx := context.Background()
	now := time.Now()
	openID := string(msg.FromUserName)
	err := o.userRepository.UserUnSubscribe(ctx, &now, openID)
	if err != nil {
		logger.Error(fmt.Sprintf("unsubscribe official account err: %+v", err))
	}
	return nil
}

// AutoReply ...
func (o *OfficialAccountService) AutoReply(msg message.MixMessage) *message.Reply {
	// TODO
	return nil
}
