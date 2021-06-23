package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

/*
SMTP rfc https://www.ietf.org/rfc/rfc2554.txt
*/

const (
	ContentTypeHTML = "Content-Type: text/html;charset=UTF-8"
)

func main() {
	user := "lixd@mail.lixueduan.com"
	nickName := "昵称"
	password := "qaz12345"
	host := "localhost:25"
	to := "xueduan.li@gmail.com"
	subject := "邮件标记"
	body := `empty body`
	fmt.Println("send email")
	err := SendByNative(nickName, user, password, host, to, subject, body)
	if err != nil {
		fmt.Println("Send mail error:", err)
		return
	}
	fmt.Println("Send mail success!")
}

// SendByNative 使用原始库发送
func SendByNative(nickname, user, password, host, to, subject, body string) error {
	/*
		From:意琦行<admin@lixueduan.com>
		To:xueduanli@163.com
		Subject:注册成功
		Content-Type: text/html;charset=UTF-8

		{there is message body}
	*/
	s := fmt.Sprintf("From:%s<%s>\nTo:%s\r\nSubject:%s\r\n%s\r\n\r\n%s",
		nickname, user, to, subject, ContentTypeHTML, body)
	msg := []byte(s)
	sendTo := strings.Split(to, ";")
	auth := smtp.PlainAuth("", "admin", "root", host)
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}
