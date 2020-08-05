package geektime

import (
	"fmt"
	"testing"
)

const Number = 24

func TestLevel1(t *testing.T) {
	for i := 0; i < Number; i++ {
		fmt.Println(Level1(i))
	}
}
func TestLevel2(t *testing.T) {
	for i := 0; i < Number; i++ {
		fmt.Println(Level2(i))
	}
}
func TestLevel3(t *testing.T) {
	for i := 0; i < Number; i++ {
		fmt.Println(Level3(i))
	}
}
func TestLevel4(t *testing.T) {
	for i := 0; i < Number; i++ {
		fmt.Println(Level4(i))
	}
}

//  39611 ns/op
func BenchmarkLevel1(b *testing.B) {
	for i := 0; i < 1; i++ {
		Level1(Number)
	}
}

// 11.1 ns/op
func BenchmarkLevel2(b *testing.B) {
	for i := 0; i < 1; i++ {
		Level2(Number)
	}
}

func BenchmarkLevel3(b *testing.B) {
	for i := 0; i < 1; i++ {
		Level3(Number)
	}
}
func BenchmarkLevel4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Level4(Number)
	}
}
