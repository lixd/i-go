package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"
)

/*
1) 每增加一个子协程就把其对应的函数地址存放到 _p_.runnext 而把 _p_.runnext 原来的地址（即上一个子协程对应的函数地址）移动到队列 _p_.runq 里面;
	所以 for 循环结束后 _p_.runnext 存放的就是最后一个子协程对应的函数地址（即输出 9的那个子协程）
2) 当开始执行子协程对应的函数时，首先执行_p_.runnext 对应的函数，然后按先进先出的顺序执行队列_p_.runq 里的函数。
*/

func main() {

	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			fmt.Print(idx)
		}(i)
		// if i%2 != 0 {
		// 	time.Sleep(time.Microsecond)
		// }
	}
	// Out: 1032547698
	time.Sleep(time.Second * 2)
	return

	score := 70.007
	score -= 100
	score += 0.001
	fmt.Println("score:", score)
	fmt.Println("Floor:", math.Floor(score))
	fmt.Println("add:", score+math.Abs(math.Floor(score)))
	fmt.Println("add:", Decimal3(score+math.Abs(math.Floor(score))))
}

func Decimal3(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
	return value
}
