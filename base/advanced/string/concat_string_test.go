package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 3

// Sprintf
func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < numbers; j++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
}

// strings.Builder
func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
}

// bytes.Buffer
func BenchmarkBytesBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j := 0; j < numbers; j++ {
			buf.WriteString(strconv.Itoa(j))
		}
		_ = buf.String()
	}
}

// Add
func BenchmarkStringAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < numbers; j++ {
			s += strconv.Itoa(i)
		}
	}
}
