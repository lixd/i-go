package aes

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAESECBEncrypt(t *testing.T) {
	key := "a1b2c3d4"
	fmt.Println("PadKey", string(generateKey([]byte(key))))

	plaintext := "2b0231214e9e48ffab777564f5f2a890"

	// 加密
	encryptedText := AESECBEncrypt([]byte(plaintext), []byte(key))
	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedText)
	fmt.Println("加密后的文本:", encryptedBase64)

	decodeString, _ := base64.StdEncoding.DecodeString(encryptedBase64)
	decryptedText := AESECBDecrypt(decodeString, []byte(key))

	// 输出解密后的文本
	fmt.Println("解密后的文本:", string(decryptedText))
}
