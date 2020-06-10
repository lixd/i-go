package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		/*		defer func(v int) {
				fmt.Println(v)
			}(i)*/
		defer func() {
			fmt.Println(i + 1)
		}()
	}
}
