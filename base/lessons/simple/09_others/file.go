package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 打开文件
	// file -->file文件对象 file指针 file文件句柄几种叫法
	file, e := os.Open("d:/test.txt")
	if e != nil {
		fmt.Println("open file error=", e)
	}
	// &{0xc000084780} 可以看出 文件就是 一个指针
	// fmt.Println(file)

	// 关闭文件
	defer file.Close()

	/*
		const (
			defaultBufSize = 4096
		)
	*/
	// 创建一个reader 带缓冲的 默认缓冲大小为 4096
	reader := bufio.NewReader(file)

	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		// 输出内容
		fmt.Print(str)
		if err == io.EOF { // io.EOF表示读到文件的末尾
			break
		}

	}
	fmt.Println()
	fmt.Println("文件读取结束")

	fmt.Println("----------------")
	// 使用 ioutil.ReadFile 一次性读取文件
	bytes, e := ioutil.ReadFile("d:/test.txt")
	if e != nil {
		fmt.Println("open file error=", e)
	}
	fmt.Println(string(bytes))

	fmt.Println("----------------")

	openFile, e := os.OpenFile("d:/test2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("OpenFile error=", e)
	}
	defer openFile.Close()
	writer := bufio.NewWriter(openFile)
	for i := 0; i < 5; i++ {
		// 调用WriteString 时先写入到缓存
		// writer.Flush()后才真正将缓冲的数据写到磁盘
		writer.WriteString("NewString \n")
	}
	e = writer.Flush()
	if e != nil {
		fmt.Println("Flush error=", e)
	}

	fmt.Println("-------------")
	src := "d:/abc.png"
	dst := "d:/def.png"
	written, e := CopyFile(src, dst)
	if e != nil {
		fmt.Printf("Copy error: %v \n", e)
	} else {
		fmt.Println("Copy Success")
	}
	fmt.Println(written)
	fmt.Println("-------------")
	fmt.Println("命令行的所有参数", len(os.Args))
	for i, value := range os.Args {
		fmt.Printf("index:%d value:%v \n", i, value)
	}

	// flag 解析命令行参数
	// 定义几个变量 用于接收命令行参数

	var user string
	var pwd string
	var host string
	var port int
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

func CopyFile(src string, dst string) (written int64, err error) {
	srcFile, e := os.Open(src)
	if e != nil {
		fmt.Printf("Open error: %v \n", e)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, e := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		fmt.Printf("Open error: %v \n", e)
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}
