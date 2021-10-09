package crc32

import (
	"fmt"
	"testing"
)

const URL = "https://www.lixueduan.com"

func Test_HashCRC32(t *testing.T) {
	crc32 := HashCRC32(URL)
	fmt.Println(crc32)
}

func Test_HashCRC64(t *testing.T) {
	fmt.Println(HashCRC64(URL))
}
