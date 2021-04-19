package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// 定义几个变量 用于接收命令行参数
	user string
	pwd  string
	host string
	port int
)

func main() {
	args()
	//	cmd()
}

func args() {
	arguments := os.Args
	// 第一个参数为默认参数 一般是可执行文件位置
	if len(arguments) == 1 {
		fmt.Printf("usage:permissions filename\n")
		fmt.Println(arguments)
		return
	}
}

func cmd() {
	// flag 解析命令行参数
	// 参数一: &user 用来接收用户命令行中输入的参数的值
	// 参数二："u" 就是 -u  指定是这个参数
	// 参数三："" 默认值 这里为空字符串
	// 参数四："用户名默认为空" 说明文字
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "p", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名默认为localhost")
	flag.IntVar(&port, "P", 3306, "端口号默认为3306")
	// 这里有个非常重要的操作 Parse 必须调用该方法
	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。
	// 未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	fmt.Printf("user:%v,pwd:%v,host:%v,port:%v", user, pwd, host, port)
}
