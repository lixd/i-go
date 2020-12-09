package hash

import (
	"github.com/spaolacci/murmur3"
)

/*
bloom过滤hash工具类
根据数组大小和hash函数个数 计算出多个hash
用于对应一个值 降低冲突
*/

type BloomFilter struct {
	m uint // bitmap大小
	k uint // hash函数个数 越大冲突越小
}

func NewBloomFilterHash(m, k uint) *BloomFilter {
	if m <= 0 {
		m = 1
	}
	if k <= 0 {
		k = 1
	}
	return &BloomFilter{
		m: m,
		k: k,
	}
}

// BloomHash 根据设置的m和k值 计算出多个hash值
/*
BenchmarkNewBloomFilterHash-4   	 6874375	       170 ns/op
*/
func (bf *BloomFilter) BloomHash(data []byte) []uint64 {
	if len(data) == 0 {
		return []uint64{0}
	}
	var bits = make([]uint64, 0, bf.k)
	hashes := baseHashes(data)
	for i := uint(0); i < bf.k; i++ {
		bit := bf.Location(hashes, i)
		bits = append(bits, bit)
	}
	return bits
}

// location returns the ith hashed location using the four base hash values
func (bf *BloomFilter) Location(h [4]uint64, i uint) uint64 {
	return location(h, i) % uint64(bf.m)
}

// location 根据base hash和传入的位置i来确定具体某一位的hash值
func location(h [4]uint64, i uint) uint64 {
	ii := uint64(i)
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}

// baseHashes murmur3 hash 计算出4个base hash值
func baseHashes(data []byte) [4]uint64 {
	a1 := []byte{1} // to grab another bit of data
	her := murmur3.New128()
	_, _ = her.Write(data) // #nosec
	v1, v2 := her.Sum128()
	_, _ = her.Write(a1) // #nosec
	v3, v4 := her.Sum128()
	return [4]uint64{
		v1, v2, v3, v4,
	}
}
