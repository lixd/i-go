package main

import (
	"testing"

	"github.com/willf/bloom"
)

func BenchmarkBloom(b *testing.B) {
	filter := bloom.New(20000000, 5)
	filter.Add([]byte("Golang")) // 添加数据
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = filter.Test([]byte("Golang")) // 测试是否存在
	}
}
