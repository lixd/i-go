package main

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"i-go/utils"
	"i-go/utils/hash"
)

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

	bloomHash := rbf.bf.BloomHash(data)
	cmders, err := rbf.rc.Pipelined(func(pipeLiner redis.Pipeliner) error {
		for _, v := range bloomHash {
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
		if v.(*redis.IntCmd).Val() != 1 {
			return false
		}
	}
	return true
}
