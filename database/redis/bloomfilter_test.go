package main

import (
	"github.com/sirupsen/logrus"
	"i-go/core/db/redisdb"
	"testing"
)

func TestRedisBloomFilter(t *testing.T) {
	var (
		key  = "firstKey"
		data = []byte("bloomFilter")
	)
	rc := redisdb.NewConn()
	bf := NewBloomFilter(1000*20, 3, rc)
	bf.Set(key, data)
	isContains := bf.isContains(key, data)
	logrus.Infof("res:%v", isContains)
}

func BenchmarkBloomFilter(b *testing.B) {
	Init("D:/lillusory/projects/i-go/conf/config.json")
	var (
		key  = "firstKey"
		data = []byte("bloomFilter")
	)
	for i := 0; i < b.N; i++ {
		bf := NewBloomFilter(1000*20, 3, rc)
		bf.Set(key, data)
		_ = bf.isContains(key, data)
	}
}
func BenchmarkBloomFilter2(b *testing.B) {
	Init("D:/lillusory/projects/i-go/conf/config.json")
	rc := redisdb.NewConn()
	for i := 0; i < b.N; i++ {
		rc.Get("user")
	}
}
