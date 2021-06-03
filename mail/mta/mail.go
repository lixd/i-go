package mta

import "fmt"

// From 发件人
type From struct {
	Nickname string // 昵称
	Address  string // 邮箱地址
}

const (
	ContentTypeHTML = "Content-Type: text/html;charset=utf-8"
)

// BuildMailData 组装邮件内容,大致格式如下:
/*
	From:意琦行<admin@lixueduan.com>
	To:xueduanli@gmail.com
	Subject:邮件标记
	Content-Type: text/html;charset=UTF-8

	{邮件内容}
*/
func BuildMailData(from From, to, subject, contentType, body string) []byte {
	data := fmt.Sprintf("From:%s<%s>\nTo:%s\r\nSubject:%s\r\n%s\r\n\r\n%s",
		from.Nickname, from.Address, to, subject, contentType, body)
	return []byte(data)
}

// ExampleMail 邮件内容测试
func ExampleMail() []byte {
	from := From{
		Nickname: "意琦行",
		Address:  "admin@lixueduan.com",
	}
	to := "xx@gmail.com"
	subject := "SMTP Service Extension for Authentication"
	// HTML 格式邮件内容
	contentType := ContentTypeHTML
	body := `
		<html>
		<body>
		<h3>
		This document specifies an Internet standards track protocol for the
	   Internet community, and requests discussion and suggestions for
	   improvements.  Please refer to the current edition of the "Internet
	   Official Protocol Standards" (STD 1) for the standardization state
	   and status of this protocol.  Distribution of this memo is unlimited.
		</h3>
		</body>
		</html>
		`
	return BuildMailData(from, to, subject, contentType, body)
}
