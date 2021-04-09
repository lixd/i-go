package convert

import (
	"fmt"
	"testing"
)

func TestBytes2String(t *testing.T) {
	str := "abcdefg"
	bytes := String2Bytes(str)
	bytes2String := Bytes2String(bytes)
	fmt.Println(str == bytes2String)
}

func BenchmarkBytes2String(b *testing.B) {
	str := "abcdefg"
	bytes := String2Bytes(str)
	for i := 0; i < b.N; i++ {
		_ = string(bytes) // 4.46 ns/op
		// _ = Bytes2String(bytes) // 0.514 ns/op
	}
}

func BenchmarkString2Bytes(b *testing.B) {
	str := "abcdefg"
	for i := 0; i < b.N; i++ {
		// _ = String2Bytes(str) // 0.520 ns/op
		_ = []byte(str) // 5.41 ns/op
	}
}
