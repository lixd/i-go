package simple

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	a, b := 1, 2
	fmt.Printf("a %d b %d \n", a, b)
	c := sum(a, b)
	fmt.Println(c)
	// letter := processLetter("abcdefghijklmn")
	letter := stringToCase("abcdefghijklmn", processLetter)
	fmt.Printf(letter)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filter := Filter(arr, isEven)
	fmt.Println(filter)
	fmt.Println("-------------------------")
	// 匿名函数 无返回值
	func(str string) {
		fmt.Printf(str)
	}("hello world")
	// 匿名函数 有返回值
	result := func(data float64) float64 {
		return math.Sqrt(data)
	}(25)
	fmt.Printf("result %.2f", result)
	// 3. 将匿名函数赋值给变量 需要时在调用
	myfunc := func(data string) string {
		return data
	}
	m := myfunc("hello func")
	fmt.Printf("m %v", m)
	sqrt()
	adder := Adder()
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d =%d \n", i, adder(i))
		fmt.Printf("fibonacci i =%d v=%d \n", i, f())
	}
}
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func Adder() func(int) int {
	// sum 为函数Adder的局部变量
	// return 的不光是一个函数 还有函数对sum的引用关系
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func sum(a, b int) (c int) {
	c = a + b
	return c
}

// 处理字符串 实现大小写交替显示
func processLetter(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

// 函数作为参数
func stringToCase(str string, mufunc func(string) string) string {
	return processLetter(str)
}

// 函数作为参数 自定义type
type caseFunc func(string) string

func stringToCase2(str string, myfunc caseFunc) string {
	return processLetter(str)
}

type myFunc func(int) bool

// 判断是否为偶数
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// 判断是否为奇数
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

// 过滤切片
func Filter(arr []int, f myFunc) []int {
	var result []int
	for _, value := range arr {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
func sqrt() {
	arr := []float64{1, 4, 9, 16, 25}
	result := 0.0
	for _, value := range arr {
		result = math.Sqrt(value)
		fmt.Printf("result %f", result)
	}
}
