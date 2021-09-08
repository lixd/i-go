package murmur

import (
	"fmt"
	"testing"
)

func TestMurmur3(t *testing.T) {
	key := []byte("hello world!")
	murmur3 := Murmur3(key)
	fmt.Println(murmur3)
}
