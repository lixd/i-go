package mta

import (
	"fmt"
	"testing"
)

func TestSender_Send(t *testing.T) {
	// 1.构建邮件内容
	from := From{
		Nickname: "意琦行",
		Address:  "lixd@tbycq.com",
	}
	// toList := []string{"xueduanli@163.com", "1033256636@qq.com", "xueduan.li@gmail.com"}
	toList := []string{"xueduan.li@gmail.com"}
	subject := "gRPC NameResolver"
	// HTML 格式邮件内容
	contentType := ContentTypeHTML
	body := `
		<html>
		<body>
		<h3>
		gRPC 中的默认 name-system 是 DNS，同时在客户端以插件形式提供了自定义 name-system 的机制。

gRPC NameResolver 会根据 name-system 选择对应的解析器，用以解析用户提供的服务器名，最后返回具体地址列表（IP+端口号）。

例如：默认使用 DNS name-system，我们只需要提供服务器的域名即端口号，NameResolver 就会使用 DNS 解析出域名对应的 IP 列表并返回。
————————————————
版权声明：本文为CSDN博主「指月小筑」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/java_1996/article/details/117190739
		</h3>
		</body>
		</html>
		`
	for _, to := range toList {
		email := BuildMailData(from, to, subject, contentType, body)
		// 	2. DKIM 加密
		err := DKIM(&email)
		if err != nil {
			fmt.Printf("DKIM to:%v error:%v\n", to, err)
			return
		}
		s := &Sender{
			Hostname: "tbycq.com",
			Insecure: true,
		}
		// 3.发送
		err = s.Send(from.Address, to, email)
		if err != nil {
			fmt.Printf("DKIM to:%v error:%v\n", to, err)
			return
		}
	}
}

// https://support.google.com/mail/answer/10336?p=NotAuthorizedError&visit_id=637582973691490016-1821529943&rd=1
