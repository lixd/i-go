package proxyutil

import (
	"fmt"
	"testing"
)

func Test_randUA(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(randUA())
	}
}
