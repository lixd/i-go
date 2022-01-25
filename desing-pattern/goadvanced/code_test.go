package goadvanced

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	ch := Random(10)
	for v := range ch {
		fmt.Print(v)
	}
}
