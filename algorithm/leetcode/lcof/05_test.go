package lcof

import (
	"testing"
)

func Test_replaceSpace(t *testing.T) {
	s := "We are happy."
	replaceSpace(s)
}

func BenchmarkReplaceSpace(b *testing.B) {
	s := "We are happy."
	for i := 0; i < b.N; i++ {
		replaceSpace(s)
	}
}
func BenchmarkReplaceSpace2(b *testing.B) {
	s := "We are happy."
	for i := 0; i < b.N; i++ {
		replaceSpaceStand(s)
	}
}
func BenchmarkReplaceSpace3(b *testing.B) {
	s := "We are happy."
	for i := 0; i < b.N; i++ {
		replaceSpace2(s)
	}
}
func BenchmarkReplaceSpace4(b *testing.B) {
	s := "We are happy."
	for i := 0; i < b.N; i++ {
		replaceSpace3(s)
	}
}
