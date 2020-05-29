package main

import (
	"fmt"
	"time"
)

func t() {
	add := time.Now().Add(time.Hour * 24)
	date := time.Now().AddDate(0, 0, 1)
	fmt.Println(add, date)
}

func main() {
	t()
	//simple()
}

func simple() {
	// 1.获取当前时间
	now := time.Now()
	fmt.Println(now)
	// 2.通过now获取其他日期信息
	fmt.Printf("年 %v \n", now.Year())
	fmt.Printf("月 %v \n", now.Month())
	fmt.Printf("日 %v \n", now.Day())
	fmt.Printf("时 %v \n", now.Hour())
	fmt.Printf("分 %v \n", now.Minute())
	fmt.Printf("秒 %v \n", now.Second())
	// 3.格式化日期时间
	// 方式一 Printf或者Sprintf
	fmt.Printf("当前时间 %02d-%02d-%02d %02d:%02d:%02d \n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	dataStr := fmt.Sprintf("当前时间 %02d-%02d-%02d %02d:%02d:%02d \n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println(dataStr)
	// 方式二 time.Format()
	// 其中2006/01/02 15:04:05是固定的 类似于 yyyy-mm-dd hh:mm:ss 只能写这个 可以组合
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	// 2019/05/19 13:53:27
	fmt.Println()
	fmt.Printf(now.Format("2006/01/02"))
	// 2019/05/19
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	// 13:55:02
	fmt.Println()
	// 5.休眠
	// 需求1：每秒打印一个数 100时退出
	// 需求2：每0.1秒打印一个数 100时退出
	// for i := 0; i <= 100; i++ {
	// 	fmt.Println(i)
	// 	//time.Sleep(time.Second) 正确
	// 	//time.Sleep(time.Second*0.1)//错误 无法通过编译
	// 	time.Sleep(time.Millisecond*100)
	// }
	// 6.Unix和UnixNano
	fmt.Printf("Unix时间戳 %v,UnixNano时间戳 %v \n", now.Unix(), now.UnixNano())
}
