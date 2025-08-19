package service

import (
	"bookkeeping/internal/model"
	"bookkeeping/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"time"

	redisPlugin "github.com/Looper56/plugin/redis"
	"github.com/go-redis/redis/v8"
)

const (
	SessionTTL         = 3600 * 72 * time.Second
	SessionCachePrefix = "bookkeeping:user_cache:"
)

// Token ...
type Token struct {
	Token string `json:"token"`
}

// SessionService session service
type SessionService struct {
	redisPlugin.Connector
}

// NewSessionService init
func NewSessionService() *SessionService {
	return &SessionService{}
}

// SaveSession ...
func (s *SessionService) SaveSession(ctx context.Context, session *model.Session) (*Token, error) {
	var err error
	var sessionID string

	for i := 0; i < 2; i++ {
		var sessionBytes []byte
		sessionBytes, err = json.Marshal(session)
		if err != nil {
			break
		}
		sessionID = s.generateSessionKey(session.OpenID)
		sessionKey := fmt.Sprintf("%s%s", SessionCachePrefix, sessionID)
		res := s.Redis().SetNX(ctx, sessionKey, string(sessionBytes), SessionTTL)
		err = res.Err()
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return &Token{Token: sessionID}, nil
}

// UpdateSession ...
func (s *SessionService) UpdateSession(ctx context.Context, sessionID string, session *model.Session) error {
	sessionBytes, err := json.Marshal(session)
	if err != nil {
		return err
	}
	sessionKey := fmt.Sprintf("%s%s", SessionCachePrefix, sessionID)
	res := s.Redis().SetXX(ctx, sessionKey, string(sessionBytes), SessionTTL)
	err = res.Err()
	return err
}

// GetSession ...
func (s *SessionService) GetSession(ctx context.Context, sessionID string) (*model.Session, error) {
	sessionKey := fmt.Sprintf("%s%s", SessionCachePrefix, sessionID)
	sessionVal, err := s.Redis().Get(ctx, sessionKey).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var session model.Session
	err = json.Unmarshal([]byte(sessionVal), &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// generateSessionKey ...
func (s *SessionService) generateSessionKey(uid string) string {
	randStr := util.RandStr(6)
	return util.MD5(fmt.Sprintf("%s%s", uid, randStr))
}
