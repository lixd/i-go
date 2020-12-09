package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

/*
使用 bcrypt 进行密码加密
加密后的格式一般为：

$2a$10$IlMH6aIithKwEPhs3.wJ1.HQyyTMf1Sq7cKPluB7vvPWGwh/Oi5vK

其中: $是分割符，无意义；2a是bcrypt加密版本号；10是cost的值；而后的前22位是salt值；再然后的字符串就是密码的密文了。

*/
func main() {
	bcryptTest()
}

func bcryptTest() {
	password := "61f03b0dd25e46a7"
	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("密文:", string(hash))
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("compare ok")
	}
}
