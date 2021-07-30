package proxyutil

import (
	"time"

	"i-go/core/db/redisdb"
)

const (
	KeyProxyServer      = "z:push:proxy:ip"
	KeyProxyServerSocks = "z:push:proxy:ip:socks"
)

// Len 集合长度
func Len(key string) int64 {
	return redisdb.Cli.SCard(key).Val()
}

// Add 写入集合
func Add(ps []string, key string) {
	if len(ps) == 0 {
		return
	}
	redisdb.Cli.SAdd(key, ps)
	redisdb.Cli.Expire(key, time.Hour*1)
}

// Pop 从集合在获取并移除元素
func Pop(key string) string {
	val := redisdb.Cli.SRandMember(key).Val()
	redisdb.Cli.SRem(key, val)
	return val
}
