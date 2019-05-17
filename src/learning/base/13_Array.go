package main

import "fmt"

//声明数组1
//未初始化有默认值 根据类型不同而不同
var num1 [5]int
var num2 = [5]int{1, 2, 3, 4, 5}

func main() {
	//声明数组2
	a := [4]float64{1.1, 2.2, 3.3, 4.4444}
	fmt.Println(a)

	b := [...]int{2, 3, 4, 5, 6}
	fmt.Println(b)

	//声明遍历1
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	//声明遍历2
	for _, value := range b {
		fmt.Println(value)
	}
	fmt.Printf("未初始化数组 %T %v \n", num1, num1)

	for _, value := range num1 {
		fmt.Println(value)

	}
	c := [3][4]int{{1, 3, 5, 7}, {2, 4, 6, 8}, {3, 6, 9, 12}}
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			fmt.Printf("c[%d][%d]=%d\n", i, j, c[i][j])
		}

	}
}
