package lock

import (
	"i-go/core/db/redisdb"
	"time"
)

type redisLock struct {
}

var RedisLock = &redisLock{}

const (
	// lua脚本保证原子性
	ReleaseLockLua = `
				if redis.call("get",KEYS[1]) == ARGV[1] then
					return redis.call("del",KEYS[1])
				else
					return 0
				end
`
)

var rc = redisdb.RedisClient

// GetLock 获取锁 增加随机值防止误释放锁
func (redisLock) GetLock(lockName, randomValue string, expire time.Duration) bool {
	nx := rc.SetNX(lockName, randomValue, expire)
	return nx.Val()
}

// ReleaseLock 释放锁 释放时要检测和获取时是相同的随机值才能释放锁
func (redisLock) ReleaseLock(lockName, randomValue string) {
	_ = rc.Eval(ReleaseLockLua, []string{lockName}, []string{randomValue})
}
