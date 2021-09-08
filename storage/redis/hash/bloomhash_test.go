package hash

import (
	"fmt"
	"testing"

	"i-go/utils/murmur"
)

func TestBloomFilter_BloomHash(t *testing.T) {
	hash := NewBloomFilterHash(1000*10, 3)
	key := []byte("second")
	bloomHash := hash.BloomHash(key)
	fmt.Printf("res :%v\n", bloomHash)
	fmt.Printf("res2 :%v\n", murmur.Murmur3(key))
}

func BenchmarkNewBloomFilterHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash := NewBloomFilterHash(1000*20, 3)
		key := []byte("second")
		_ = hash.BloomHash(key)
	}
}

func TestA(t *testing.T) {
	type tmp struct {
	}
	type Arr struct {
		A []tmp
	}
	var a Arr
	a.A = append(a.A, tmp{})
	fmt.Println(a.A)
}
