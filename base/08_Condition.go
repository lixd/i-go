package main

import "fmt"

func main() {
	//switch fallthrough
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("A")
		fallthrough
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")
	default:
		fmt.Printf("default")
	}
	//for
	for i := 1; i < 10; i++ {
		fmt.Printf("%d \n", i)
	}
	//for
	result := 0
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			result += i
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("%d \n", result)
	//for
	count := 0
	for i := 32.0; i >= 4; i -= 1.5 {
		count++
	}
	fmt.Printf("count: %d \n", count)
	//遍历字符串
	str := "abcdefghiABCDEFGHI一丁丂"
	for i, value := range str {
		fmt.Printf("第%d位的字符是：%c \n", i, value)

	}
	//直角三角形
	for i := 0; i < 10; i++ {
		for k := 10; k > i; k-- {
			fmt.Print(" ")
		}
		for j := 10 - i; j < 10; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	//等腰三角形
	for i := 0; i <= 11; i++ {
		for x := 0; x < 11-i; x++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j += 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//1-50中不带4的数 continue
	for i := 1; i <= 50; i++ {
		if i%10 == 4 || i/10%10 == 4 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	//1-50中不带4的数 goto
	fmt.Println()
	i := 0
LOOP:
	for i < 50 {
		i++
		if i%10 == 4 || i/10%10 == 4 {
			goto LOOP
		}
		fmt.Printf("%d ", i)

	}
}
