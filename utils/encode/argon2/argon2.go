package main

import (
	"fmt"
	"golang.org/x/crypto/argon2"
)

// 首选 Argon2 次选 Scrypt 最后为 BCrypt
/*
使用 argon2 进行密码加密
*/
func main() {
	argon()
}

func argon() {
	pwd := "admin"
	key := argon2.IDKey([]byte(pwd), []byte("xxx"), 1, 64*1024, 4, 32)
	fmt.Println(string(key))
	key2 := argon2.IDKey([]byte(pwd), []byte("xxxy"), 1, 64*1024, 4, 32)
	fmt.Println(string(key2))
}
