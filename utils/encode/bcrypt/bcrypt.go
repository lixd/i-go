package bcrypt

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
)

/*
使用 bcrypt 进行密码加密
加密后的格式一般为：

$2a$10$IlMH6aIithKwEPhs3.wJ1.HQyyTMf1Sq7cKPluB7vvPWGwh/Oi5vK

其中: $是分割符，无意义；2a是bcrypt加密版本号；10是cost的值；而后的前22位是salt值；再然后的字符串就是密码的密文了。

*/
func main() {
	t()
	//Transfer()
}

func t() {
	password := "61f03b0dd25e46a7"
	passwordOK := "admin"
	passwordERR := "adminxx"
	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePW := string(hash)
	// 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println("加密后密码", encodePW)
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
	if err != nil {
		fmt.Println("pw wrong: ", err)
	} else {
		fmt.Println("pw ok")
	}
	// 错误密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordERR))
	if err != nil {
		fmt.Println("pw wrong: ", err)
	} else {
		fmt.Println("pw ok")
	}
	unencodedSalt := make([]byte, 10)
	_, err = io.ReadFull(rand.Reader, unencodedSalt)
	if err != nil {
		fmt.Println("read: ", err)
	}
	fmt.Println(unencodedSalt)
	err = bcrypt.CompareHashAndPassword([]byte("$2a$10$irBhoZJYVO3BuFI9ELrRKOYoLGeAUqaSSDWglismOz3Vwcjbj.fXS"), []byte("61f03b0dd25e46a7"))
	if err != nil {
		fmt.Println("test wrong: ", err)
	} else {
		fmt.Println("test ok")
	}
}

func Transfer() {
	m := make(map[string]string)

	m["5dd3f51c27d"] = ""
	m["c026473d98964f6193f0ded43203e5e4"] = ""
	m["451dade2f354a"] = ""
	m["61f03b0dd25e46a7"] = ""
	m["aa1b7b594cba8fd6"] = ""
	m["aa1b7b594cba8fd6"] = ""
	m["1b7b594cba8fd6"] = ""

	for k, _ := range m {
		bytes, err := bcrypt.GenerateFromPassword([]byte(k), bcrypt.DefaultCost)
		if err == nil {
			m[k] = string(bytes)
		}
	}
	for k, v := range m {
		fmt.Printf("k:%v,v:%v \n", k, v)
	}

}
