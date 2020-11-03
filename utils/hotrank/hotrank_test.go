package hotrank

import (
	"fmt"
	"testing"
)

func TestNewtonsLawOfCooling(t *testing.T) {
	cooling := NewtonsLawOfCooling(1, 24)
	fmt.Println(cooling)
}

/*
1) 1) (integer) 33 //每条日志的唯一ID编号
2) (integer) 1600990583 //命令执行时的时间戳
3) (integer) 20906 //命令执行的时长，单位是微秒
4)
  1) "keys" //具体的执行命令和参数
  2) "abc*"
5) "127.0.0.1:54793" //客户端的IP和端口号
6) "" //客户端的名称，此处为空

*/
