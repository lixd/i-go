package main

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"i-go/storage/redis/hash"
	"i-go/utils"
)

// 使用 Redis Bitmap 自己实现的一个 BloomFilter

type redisBloomFilter struct {
	bf *hash.BloomFilter
	rc *redis.Client
}

func NewBloomFilter(m, k uint, rc *redis.Client) *redisBloomFilter {
	defer utils.Trace("NewBloomFilter")()
	bf := hash.NewBloomFilterHash(m, k)
	return &redisBloomFilter{bf: bf,
		rc: rc}
}

// Set 将data添加到当前key中
func (rbf *redisBloomFilter) Set(key string, data []byte) {
	defer utils.Trace("Set")()
	// 将当前key根据多个hash函数计算出多个hash值
	bloomHash := rbf.bf.BloomHash(data)
	cmders, err := rbf.rc.Pipelined(func(pipeLiner redis.Pipeliner) error {
		for _, v := range bloomHash {
			// 并将对应位置置1
			pipeLiner.SetBit(key, int64(v), 1)
		}
		return nil
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"scenes": "bloom filter SetBit"}).Error(err)
	}
	logrus.Infof("pipeLine setBit res:%v", cmders)

}

// isContains 检测当前key中是否存在data
func (rbf *redisBloomFilter) isContains(key string, data []byte) bool {
	defer utils.Trace("isContains")()
	// 同样的 计算出多个hash值后看是否每位都为1
	bloomHash := rbf.bf.BloomHash(data)
	cmders, err := rbf.rc.Pipelined(func(pipeLiner redis.Pipeliner) error {
		for _, v := range bloomHash {
			pipeLiner.GetBit(key, int64(v))
		}
		return nil
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"scenes": "bloom filter GetBit error"}).Error(err)
		return false
	}
	logrus.Infof("pipeLine GetBit res:%v", cmders)
	for _, v := range cmders {
		// 这里需要转成对应类型
		// 只有有一个位上不为1则表示该值肯定不在数组中
		if v.(*redis.IntCmd).Val() != 1 {
			return false
		}
	}
	return true
}
