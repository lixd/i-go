package srt

import "testing"

func BenchmarkIsValidPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidPassword("password")
	}
}

func BenchmarkIsValidUsername(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidUsername("username")
	}
}
