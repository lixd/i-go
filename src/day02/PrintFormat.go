package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	//通用格式
	str := "illusory"
	//string illusory
	fmt.Printf("%T %v \n", str, str)

	p := point{1, 2}
	//main.point {1 2}
	fmt.Printf("%T %v \n", p, p)
	//布尔值
	fmt.Printf("%T %t \n", true, true)
	//整形
	fmt.Printf("%T %d \n", 123, 123)
	fmt.Printf("%T %5d \n", 123, 123)
	fmt.Printf("%T %05d \n", 123, 123)
	fmt.Printf("%T %b \n", 123, 123)
	//进制转换
	sprintf := fmt.Sprintf("%b", 123)
	fmt.Printf("%T %v \n", sprintf, sprintf)
	fmt.Printf("%x \n", 123)
	fmt.Printf("%X \n", 123)
	fmt.Printf("%U \n", 123)
	//浮点型
	//123.123000
	fmt.Printf("%f \n", 123.123)
	//123.13 会四舍五入
	fmt.Printf("%.2f \n", 123.1251231)
	//字符串
	//幻境云图
	fmt.Printf("%s \n", "幻境云图")
	//"幻境云图"  加上双引号
	fmt.Printf("%q \n", "幻境云图")

}
