package mta

import (
	"testing"
)

func TestDKIM(t *testing.T) {
	email := ExampleMail()
	err := DKIM(&email)
	if err != nil {
		t.Fatal(err)
	}
}

// 2892294 ns/op 2.8ms
func BenchmarkDKIM(b *testing.B) {
	email := ExampleMail()
	for i := 0; i < b.N; i++ {
		_ = DKIM(&email)
	}
}
