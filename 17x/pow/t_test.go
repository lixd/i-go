package main

import (
	"fmt"
	"testing"
)

func Test_isOk(t *testing.T) {
	type args struct {
		str   string
		com   string
		level int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{
			str:   "012345",
			com:   "02",
			level: 1,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOk(tt.args.str, tt.args.com, tt.args.level); got != tt.want {
				t.Errorf("isOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashSha256(t *testing.T) {
	fmt.Println(1 > 0 && 2 < 1 || 1 > 0 && 2 > 1 && 3 > 2)
	return
	str := "vid" + "4729934"
	sha256 := hashSha256([]byte(str))
	fmt.Printf("src:%s hash:%s\n", str, sha256)
}

// 358 ns/op 6200ns 17
func BenchmarkHash(b *testing.B) {
	str := "vid" + "4729934"
	data := []byte(str)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hashSha256(data)
	}
}
