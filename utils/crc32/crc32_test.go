package crc32

import (
	"fmt"
	"testing"
)

const URL = "https://www.lixueduan.com"

func Test_getIntvalKey(t *testing.T) {
	fmt.Println(HashCRC32(URL))
}
func Test_getIntvalKey64(t *testing.T) {
	fmt.Println(HashCRC64(URL))
}
