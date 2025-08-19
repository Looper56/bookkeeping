package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Looper56/plugin/redis"
	redisCmd "github.com/redis/go-redis/v9"
)

// EmailStatusRepository email validate
type EmailStatusRepository struct {
	redis.Connector
}

// NewEmailStatusRepository init
func NewEmailStatusRepository() *EmailStatusRepository {
	return &EmailStatusRepository{}
}

var emailKeyFmt = "bookkeeping:member_uid:%s"

// SetEmailVerificationCode set email verify code timeliness
func (e *EmailStatusRepository) SetEmailVerificationCode(ctx context.Context, memberUID, email, code string) error {
	key := fmt.Sprintf(emailKeyFmt, email)
	redisClient := e.Redis()
	var cmd *redisCmd.IntCmd
	cmd = redisClient.HSet(ctx, key, memberUID, code)
	ok, err := redisClient.Expire(ctx, key, 300*time.Second).Result()
	if !ok {
		return err
	}
	err = cmd.Err()
	if err != nil {
		return err
	}
	return nil
}

// GetEmailVerificationCode get email verify code timeliness
func (e *EmailStatusRepository) GetEmailVerificationCode(ctx context.Context, memberUID, email string) (string, error) {
	key := fmt.Sprintf(emailKeyFmt, email)
	redisClient := e.Redis()
	code, err := redisClient.HGet(ctx, key, memberUID).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}
