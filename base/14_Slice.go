package main

import "fmt"

// 声明切片1 未初始化时为nil
var s1 []int

// 声明切片2.1 make方式声明的切片不为nil
var s2 []int = make([]int, 5)
var s3 []int = make([]int, 5, 7)

// 方式3 直接声明切片并赋值
var s6 []int = []int{1, 3, 4, 5, 7}

func main() {
	slice2()
}

func slice2() {
	// 基本语法：var 切片名 []type = make([], len, [cap]);参数说明：type是数据类型、len是大小、cap是切片容量（容量必须>=长度）
	sl1 := make([]int, 10, 10)
	sl2 := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		sl1 = append(sl1, i)
	}
	i := copy(sl2, sl1)
	fmt.Println("copy Number: ", i)
	printMsg(sl1)
}

func slice1() {
	s4 := make([]int, 5)
	s5 := make([]int, 5, 7)
	printMsg(s1)
	printMsg(s2)
	printMsg(s3)
	printMsg(s4)
	printMsg(s5)
	printMsg(s6)
	// 切片遍历 普通for循环
	for i := 0; i < len(s6); i++ {
		fmt.Println(s6[i])
	}
	// 切片遍历 for range
	fmt.Println("------------")
	for _, value := range s6 {
		fmt.Println(value)
	}
	// append 动态扩容
	ints := append(s6, 100, 200, 300)
	// 返回的即是扩容后的切片
	fmt.Println(ints)
	fmt.Println(s6)
	fmt.Println(s4)
	copy(s4, s6)
	fmt.Println(s4)
	var s11 []int = []int{1, 3, 4, 5, 7}
	s12 := make([]int, 1)
	// 全是默认值 {0,0,0,0,0}
	copy(s12, s11)
	// 将s1拷贝到s2 由于s2长度为1 所以最终只会拷贝1个元素 但是不会报错
	fmt.Println(s12)
	// string底层是byte数组，因此string也可以进行切片处理
	str := "hello go"
	slices := str[6:]
	fmt.Println(slices)
	// go
	arr1 := []byte(str)
	arr1[6] = 'G'
	str = string(arr1)
	fmt.Println(str)
	arr2 := []rune(str)
	arr2[6] = '白'
	str = string(arr2)
	fmt.Println(str)
	names := [4]string{"Java", "C", "C++", "Go"}
	lang := ""
	fmt.Scanln(&lang)
	index := -1
	for i, value := range names {
		if value == lang {
			index = i
		}
	}
	if index != -1 {
		fmt.Printf("您输入的%s 存在\n", lang)
	} else {
		fmt.Println("不存在")
	}
}
func printMsg(s []int) {
	fmt.Printf("len=%d cap=%d slice=%v \n", len(s), cap(s), s)
}
