package main

import (
	"fmt"

	"golang.org/x/crypto/argon2"
)

/*
使用 argon2 进行密码加密
*/
func main() {
	argon()
}

func argon() {
	pwd := "admin"
	salt1 := "xxx"
	salt2 := "xxxx"
	key := argon2.IDKey([]byte(pwd), []byte(salt1), 1, 64*1024, 4, 32)
	fmt.Println(string(key))
	key2 := argon2.IDKey([]byte(pwd), []byte(salt2), 1, 64*1024, 4, 32)
	fmt.Println(string(key2))
	fmt.Println(string(key) == string(key2))
}
