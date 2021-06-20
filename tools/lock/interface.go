package lock

import "time"

type ILock interface {
	Lock(key string, value interface{}, expire time.Duration) bool
	UnLock(key string, value interface{}) error
}
