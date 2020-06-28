package _17x

import (
	"fmt"
)

func main() {
	/*	array1 := [3]string{"a", "b", "c"}
		fmt.Printf("The array: %v\n", array1)
		array2 := modifyArray(array1)
		fmt.Printf("The modified array: %v\n", array2)
		fmt.Printf("The original array: %v\n", array1)*/

}

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

func forRange() {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}
