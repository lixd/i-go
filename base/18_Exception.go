package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randArray()
	test()
	err := loadConfig("init.conf")
	if err != nil {
		// 如果读取错误则输出错误并终止程序
		panic(err)
	} else {
		fmt.Println("配置文件读取成功，程序继续执行")
	}
}
func test() int {
	// 使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover内置函数 捕获异常
		if err != nil {  // 如果err不为nil说明出现异常了
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	// 这里会抛出异常 runtime error: integer divide by zero
	res := num1 / num2
	return res
}

// 函数读取init.conf的配置文件
// 如果文件名传入不对就返回自定义错误
func loadConfig(name string) (err error) {
	if name == "init.conf" {
		// 正确则读取
		return nil
	} else {
		// 否则返回一个自定义错误
		return errors.New("读取文件错误。。")
	}
}

var arrs [5]int

// 随机生成数组并反转打印
func randArray() {
	rand.Seed(time.Now().UnixNano())
	len := len(arrs)
	// 随机生成数组
	for i := 0; i < len; i++ {
		arrs[i] = rand.Intn(100)
	}
	fmt.Println(arrs)
	// 反转打印
	for i := 0; i < len/2; i++ {
		arrs[i], arrs[len-i-1] = arrs[len-i-1], arrs[i]
	}
	fmt.Println(arrs)
}
