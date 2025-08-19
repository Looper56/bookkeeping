package atomiclock

import (
	"context"
	"errors"
	"time"

	redisPlugin "github.com/Looper56/plugin/redis"
	"github.com/bsm/redislock"
)

// ErrNotObtained is returned when a lock cannot be obtained.
var ErrNotObtained = errors.New("atomic locks: not obtained")

// LockObject lock object interface
type LockObject interface {
	// GetLockKey 上锁key
	GetLockKey() string
	// GetLockTTL 上说周期s
	GetLockTTL() time.Duration
}

// AtomicLock distributed atomic lock
type AtomicLock struct {
	redisPlugin.Connector
	lockClient *redislock.Client
	lock       *redislock.Lock
	Error      error
}

// NewAtomicLock init atomic lock
func NewAtomicLock() *AtomicLock {
	atomicLock := &AtomicLock{}
	atomicLock.lockClient = redislock.New(atomicLock.Redis())
	return atomicLock
}

// Obtain try lock
func (a *AtomicLock) Obtain(ctx context.Context, key string, ttl time.Duration) bool {
	lock, err := a.lockClient.Obtain(ctx, key, ttl, nil)
	if err != nil {
		if errors.Is(err, redislock.ErrNotObtained) {
			a.Error = ErrNotObtained
		} else {
			a.Error = err
		}
		return false
	}
	a.lock = lock
	return true
}

// Release release lock
func (a *AtomicLock) Release(ctx context.Context) error {
	if a.lock == nil {
		return errors.New("lock is not init")
	}
	if a.Error != nil {
		return a.Error
	}
	return a.lock.Release(ctx)
}
