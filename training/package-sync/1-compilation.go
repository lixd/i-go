package main

import "fmt"

var counter int

// go tool compile -S 1-compilation.go
/*
counter++ 并不是原子操作，底层是分3步完成:
MOVQ    "".counter(SB), AX // 1.将counter的值赋给AX
INCQ    AX				// 2. AX 自增
MOVQ    AX, "".counter(SB) // 3. 将AX的值赋给counter
*/
func main() {
	counter++
	fmt.Println(counter)
}
