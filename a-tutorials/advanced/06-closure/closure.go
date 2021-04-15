package main

import "fmt"

// func main() {
// 	f1 := incrementer()
// 	f2 := incrementer()
// 	fmt.Println(f1())
// 	fmt.Println(f2())
// }

func incrementer() func() int {
	// 局部变量 i 除了初始化之外没有被修改过，所以闭包就直接复制了一份到栈上。
	i := 2
	return func() int {
		return i
	}
}

func main() {
	fs := create()
	for i := 0; i < len(fs); i++ {
		fs[i]()
	}
}

func create() (fs [2]func()) {
	// 局部变量 i 除了初始化之外还被修改过，同时还被闭包函数捕获，因此会分配到堆上，栈上保存的都是堆上变量的地址。
	// 所以最后这个i打印出来都是一个值(因为都指向堆上变量i)。
	for i := 0; i < 2; i++ {
		fs[i] = func() {
			fmt.Println(i)
		}
	}
	return
}
