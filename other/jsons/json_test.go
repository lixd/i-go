package josns

import "testing"

func BenchmarkDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default()
	}
}

func BenchmarkSpecial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Special()
	}
}
