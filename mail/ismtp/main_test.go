package ismtp

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	ContentTypeHTML = "Content-Type: text/html;charset=UTF-8"
)

func TestSMTP_Deliver(t *testing.T) {
	s := SMTP{
		HelloDomain: "localhost",
		Dinfo:       nil,
		STSCache:    nil,
	}
	nickName := "天猫精灵"
	from := "lixd@tbycq.com"
	to := "xueduanli@163.com"
	// to := "1033256636@qq.com"
	subject := "每日一题"
	body := `
		<html>
		<body>
		<h3>
	每日精选如下:
	给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
	请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
		</h3>
		</body>
		</html>
		`
	data := fmt.Sprintf("From:%s<%s>\nTo:%s\r\nSubject:%s\r\n%s\r\n\r\n%s",
		nickName, from, to, subject, ContentTypeHTML, body)
	err, b := s.Deliver(from, to, []byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("结果:", b)
}
