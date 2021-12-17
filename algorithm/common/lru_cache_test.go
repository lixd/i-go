package common

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLRUCache(t *testing.T) {
	LRUCache := NewLRUCache(3)
	LRUCache.Set(1, 11)
	LRUCache.Set(2, 22)
	LRUCache.Set(3, 33)
	LRUCache.Set(4, 44) // 会把1移除掉
	LRUCache.Set(5, 55) // 会把1移除掉
	LRUCache.Set(6, 66) // 会把1移除掉
	for _, v := range LRUCache.m {
		logrus.Printf("key:%v value:%v", v.key, v.value)
	}
}
