package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"i-go/utils"
	"strconv"
	"strings"
	"unicode/utf8"
)

var str string = "illusory cloud 幻境云图"

// 字符串
func main() {
	//stringB()
	//forRange()
	reader()
}

func reader() {
	r := strings.NewReader(str)
	fmt.Printf("len:%v size:%v \n", r.Len(), r.Size())
	b := make([]byte, 8)
	n, err := r.Read(b)
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller()}).Error(err)
	}
	fmt.Printf("readCount:%v str:%v \n", n, string(b))
	// 读取后 len 变了 说明 len 返回的是未读取的长度
	fmt.Printf("len:%v size:%v \n", r.Len(), r.Size())

}

func stringB() {
	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
}

func forRange() {
	str := "Go爱好者"
	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}
}

func stringA() {
	// 1.len
	fmt.Printf("字符串长度 len(str)：%d \n", len(str))
	fmt.Printf("字符串个数 ：%d \n", utf8.RuneCountInString(str))
	fmt.Printf("字符串长度 转成[]rune后：%d \n", len([]rune(str)))

	// 2.遍历 处理中文问题
	str2 := []rune(str)
	for _, vlaue := range str2 {
		fmt.Printf("%c \n", vlaue)
	}
	// 3.字符串转整数
	i, err := strconv.Atoi("123A")
	if err != nil {
		fmt.Printf("转换错误：err %v \n", err)
	} else {
		fmt.Printf("转换成功 结果为：%d \n", i)
	}
	// 4.整数转字符串
	i2 := strconv.Itoa(12345)
	fmt.Printf("转换成功 结果为：%v \n", i2)
	// 5.字符串转byte
	bytes := []byte("hello go")
	fmt.Println(bytes)
	// 6.byte转字符串
	str3 := string([]byte{97, 98, 99})
	fmt.Println(str3)
	// 7.10进制转2,8,16进制
	formatInt := strconv.FormatInt(123, 2)
	fmt.Println(formatInt)

	// 8.查找子串是否在字符串中
	contains := strings.Contains("hello go", "go")
	fmt.Printf("hello go 中是否存在go %t \n", contains)
	// 9.统计一个字符串出现了几次子串
	count := strings.Count("hello go", "o")
	fmt.Printf("hello go 包含%d个o \n", count)
	// 10.不区分大小写的字符串比较
	fold := strings.EqualFold("hello", "HELLO")
	fmt.Printf("不区分大小写 HELLO hello 比较 %t \n", fold)
	// 11.返回子串在字符串中第一次出现的index值，若没有则返回-1
	index := strings.Index("hello go", "o")
	fmt.Printf("hello go中o第一次出现的索引%d \n", index)
	// 12.返回子串在字符串中最后一次出现的index值，若没有则返回-1
	lastIndex := strings.LastIndex("hello go", "o")
	fmt.Printf("hello go中o最后一次出现的索引%d \n", lastIndex)
	// 13.将指定的子串替换成另外一个子串
	replace := strings.Replace("hello go", "go", "golang", 1)
	fmt.Printf("将hello go中的go替换为golang后：%s \n", replace)
	// 14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	split := strings.Split("hello go", " ")
	fmt.Printf("将hello go按照' '拆分后：%v \n", split)
	// 15.将字符串的字母进行大小写转换
	lower := strings.ToLower("HeLLo GO")
	upper := strings.ToUpper("hello go")
	fmt.Printf("lower: %s upper %s \n", lower, upper)
	// 16.将字符串左右两边的空格去掉
	space := strings.TrimSpace(" Hello Go    ~ ")
	fmt.Printf(" Hello Go    ~ 去掉左右两边字符串 %q \n", space)
	// 17.将字符串左右两边指定的字符去掉
	trim := strings.Trim("@@Hello GO@~", "@~")
	fmt.Printf("@@Hello GO@~ 去掉左右两边的@~后 %s \n", trim)
	// 18.将字符串左边指定字符去掉
	left := strings.TrimLeft("@@Hello GO@~", "@~")
	fmt.Printf("@@Hello GO@~ 去掉左边的@~后 %s \n", left)
	// 19.将字符串右边指定字符去掉
	right := strings.TrimRight("@@Hello GO@~", "@~")
	fmt.Printf("@@Hello GO@~ 去掉右边的@~后 %s \n", right)
	// 20.判断字符串是否以指定的字符串开头
	prefix := strings.HasPrefix("hello go", "hello")
	fmt.Printf("hello go 是否以hello开头 %t \n", prefix)
	// 21.判断字符串是否以指定的字符串结束
	suffix := strings.HasSuffix("hello go", "go")
	fmt.Printf("hello go 是否以go结束 %t \n", suffix)
}
