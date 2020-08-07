package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

const (
	char = "0123456789abcdef"
)

func MD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func RandStr() (result string) {
	charArr := []byte(char)
	for i := 0; i < 4; i++ {
		result = fmt.Sprintf("%s%v", result, string(charArr[rand.Intn(16)]))
	}
	return result
}
