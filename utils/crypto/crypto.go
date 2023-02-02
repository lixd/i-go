package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

/*
NOTE: 选择加密算法时首选 Argon2 次选 Scrypt 最后为 BCrypt。

用户注册时使用 salt 对密码进行加密，DB中只存储加密后的值，不存在明文防止数据泄露导致密码也跟着泄露。
	可以未每个用户生成不同的 salt.
用户登录是，将用户提交的密码加密后再和数据库中的密文对比即可。
*/

const (
	dbPwd   = "root"
	userPwd = "root"
)

var salt = []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}

func main() {
	argon2Demo()
	scryptDemo()
	bcryptDemo()
}

func argon2Demo() {
	// dbPwd 加密,DB中就存储这个Key即可。
	key := argon2.IDKey([]byte(dbPwd), salt, 1, 64*1024, 4, 32)
	b1 := base64.StdEncoding.EncodeToString(key)
	fmt.Println("argon2: ", b1)

	// userPwd 加密

	key2 := argon2.IDKey([]byte(userPwd), salt, 1, 64*1024, 4, 32)
	b2 := base64.StdEncoding.EncodeToString(key2)
	fmt.Println("argon2: ", b2)
	// 对比二者是否一致来检测用户密码输对没有
	fmt.Println(b1 == b2)
}

func scryptDemo() {
	key, err := scrypt.Key([]byte(dbPwd), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	b1 := base64.StdEncoding.EncodeToString(key)
	fmt.Println("scrypt: ", b1)
	key2, err := scrypt.Key([]byte(userPwd), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	b2 := base64.StdEncoding.EncodeToString(key2)
	fmt.Println("scrypt: ", b2)

	fmt.Println(b1 == b2)
}

/*
使用 bcrypt 进行密码加密
加密后的格式一般为：
$2a$10$IlMH6aIithKwEPhs3.wJ1.HQyyTMf1Sq7cKPluB7vvPWGwh/Oi5vK
其中: $是分割符，无意义；2a是bcrypt加密版本号；10是cost的值；而后的前22位是salt值；再然后的字符串就是密码的密文了。
*/
func bcryptDemo() {
	hash, err := bcrypt.GenerateFromPassword([]byte(dbPwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("bcrypt: ", string(hash))

	// 密码验证 返回 error = nil 时即正确。
	err = bcrypt.CompareHashAndPassword(hash, []byte(userPwd))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}
