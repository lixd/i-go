package redis

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLFUCache(t *testing.T) {
	lfuCache := NewLFUCache(3)
	lfuCache.Set(1, 11)
	lfuCache.Set(1, 11)
	lfuCache.Set(1, 11)
	lfuCache.Set(2, 22)
	lfuCache.Set(2, 22)
	lfuCache.Set(3, 33)
	lfuCache.Set(4, 44) // 会把3移除掉
	for _, v := range lfuCache.cacheMap {
		logrus.Printf("key:%v value:%v", v.key, v.value)
	}
}
