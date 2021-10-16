package lock

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// 基于 Redis 的分布式锁

var (
	once      sync.Once
	RedisLock *redisLock
)

type redisLock struct {
	RedisCli *redis.Client
}

func NewRedisLock(cli *redis.Client) *redisLock {
	once.Do(func() {
		RedisLock = &redisLock{
			RedisCli: cli,
		}
	})
	return RedisLock
}

const (
	// lua脚本保证原子性
	unLockLua = `
				if redis.call("get",KEYS[1]) == ARGV[1] then
					return redis.call("del",KEYS[1])
				else
					return 0
				end
`
)

// Lock 获取锁 增加随机值防止误释放锁（只有值相同时才能释放锁）
func (r *redisLock) Lock(key string, value interface{}, expire time.Duration) bool {
	nx := r.RedisCli.SetNX(key, value, expire)
	return nx.Val()
}

// UnLock 释放锁 释放时要检测和获取时是相同的值才能释放锁
func (r *redisLock) UnLock(key string, value interface{}) error {
	err := r.RedisCli.Eval(unLockLua, []string{key}, value).Err()
	return err
}
