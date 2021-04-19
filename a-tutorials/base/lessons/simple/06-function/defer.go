package main

import "fmt"

// defer
/*
关键字 defer 用于注册延迟调用。
这些调用直到 return 前才被执。因此，可以用来做资源清理。
多个defer语句，按先进后出的方式执行。
defer语句中的变量，在defer声明时就决定了。
*/
//func main() {
//	defer func() {
//		fmt.Println("defer1")
//	}()
//	defer func() {
//		fmt.Println("defer2")
//	}()
//	fmt.Println("main")
//	panic("异常触发了")
//}

func main() {
	fmt.Println(doubleScore(0))    //0
	fmt.Println(doubleScore(20.0)) //40
	fmt.Println(doubleScore(50.0)) //50
}
func doubleScore(source float32) (score float32) {
	// defer 可以改变返回值
	defer func() {
		if score < 1 || score >= 100 {
			//将影响返回值
			score = source
		}
	}()
	return source + 50
}
