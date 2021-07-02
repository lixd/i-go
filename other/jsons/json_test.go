package josns

import "testing"

// 2145 ns/op
func BenchmarkDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default()
	}
}

// 899 ns/op
func BenchmarkSpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Special()
	}
}
