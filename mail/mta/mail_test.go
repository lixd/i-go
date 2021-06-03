package mta

import (
	"fmt"
	"testing"
)

func TestBuildMailData(t *testing.T) {
	from := From{
		Nickname: "意琦行",
		Address:  "admin@lixueduan.com",
	}
	to := "xueduanli@gmail.com"
	subject := "邮件标题"
	// HTML 格式邮件内容
	contentType := ContentTypeHTML
	body := `
		<html>
		<body>
		<h3>
			邮件内容
		</h3>
		</body>
		</html>
		`
	data := BuildMailData(from, to, subject, contentType, body)
	fmt.Println(data)
}
