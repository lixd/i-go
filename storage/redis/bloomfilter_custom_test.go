package main

import (
	"testing"

	"i-go/core/db/redisdb"

	"github.com/sirupsen/logrus"
)

func TestRedisBloomFilter(t *testing.T) {
	var (
		key  = "firstKey"
		data = []byte("bloomFilter")
	)
	Init("D:/lillusory/projects/i-go/conf/config.json")
	rc := redisdb.Cli
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
		bf := NewBloomFilter(1000*20, 3, redisdb.Client())
		bf.Set(key, data)
		_ = bf.isContains(key, data)
	}
}

func BenchmarkBloomFilter2(b *testing.B) {
	Init("D:/lillusory/projects/i-go/conf/config.json")
	rc := redisdb.Cli
	for i := 0; i < b.N; i++ {
		rc.Get("user")
	}
}
