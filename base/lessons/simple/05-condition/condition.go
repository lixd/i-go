package main

import "fmt"

// 流程控制
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

// switch 与 case 类型不同 无法通过编译
/*func switchA() {
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch 1 + 3 {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4 or 5 or 6")
	}
}*/

func switchB() {
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value1[4] {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4 or 5 or 6")
	}
}

// 存在相同值的 case 无法通过编译
/*func switchC() {
	value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value3[4] {
	case 0, 1, 2:
		fmt.Println("0 or 1 or 2")
	case 2, 3, 4:
		fmt.Println("2 or 3 or 4")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}
*/
// 索引表达式可以存在 相同 case 能绕过编译器检测
/*func switchD() {
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or26")
	}
}*/
