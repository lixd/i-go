package hash

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestBloomFilter_BloomHash(t *testing.T) {
	hash := NewBloomFilterHash(1000*10, 3)
	key := []byte("second")
	bloomHash := hash.BloomHash(key)
	logrus.Infof("res :%v", bloomHash)
	logrus.Infof("res2 :%v", Murmur3(key))
}

func BenchmarkNewBloomFilterHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash := NewBloomFilterHash(1000*20, 3)
		key := []byte("second")
		_ = hash.BloomHash(key)
	}
}
