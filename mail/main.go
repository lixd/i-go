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

/*
错误码
QQ: 550 Suspected bounce attacks https://service.mail.qq.com/cgi-bin/help?id=20022&no=1001602&subtype=1
疑似退信攻击
发生此问题，可能是你的邮件服务器接收了仿冒qq.com账号发出的垃圾邮件，并且你的邮件服务器没有检查发件人真实性，在某种条件下触发了退信。
qq.com发出的所有邮件均可通过SPF和DKIM校验，请根据标准直接拒绝此类仿冒qq.com的邮件。
503 身份验证
*/
