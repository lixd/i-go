package hash

import (
	"fmt"
	"strconv"
	"testing"
)

// 大致算是均匀分配
func TestNew(t *testing.T) {
	var (
		static = make(map[string]int)
	)
	m := New(100, nil)
	physical := []string{"A", "B", "C", "D", "E"}
	m.Add(physical...)
	for i := 0; i < 500_0000; i++ {
		key := strconv.Itoa(i)
		get := m.Get(key)
		static[get] += 1
	}
	fmt.Println(static)
}
