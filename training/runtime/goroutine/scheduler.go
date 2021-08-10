package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
1) 每增加一个子协程就把其对应的函数地址存放到 _p_.runnext 而把 _p_.runnext 原来的地址（即上一个子协程对应的函数地址）移动到队列 _p_.runq 里面;
	所以 for 循环结束后 _p_.runnext 存放的就是最后一个子协程对应的函数地址（即输出 9的那个子协程）
2) 当开始执行子协程对应的函数时，首先执行_p_.runnext 对应的函数，然后按先进先出的顺序执行队列_p_.runq 里的函数。
Scheduler Affinity communicate-and-wait 模式优化。Go 1.5 在 P 中引入了 runnext 特殊的一个字段，可以高优先级执行 unblock G。
*/
func main() {
	// case1()
	case2()
}

func case1() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			fmt.Print(idx)
		}(i)
	}
	// Out: 9012345678
	// 为什么每次都是9最先执行，后续按顺序执行
	/*
		1) 每增加一个子协程就把其对应的函数地址存放到 _p_.runnext 而把 _p_.runnext 原来的地址（即上一个子协程对应的函数地址）移动到队列 _p_.runq 里面;
			所以 for 循环结束后 _p_.runnext 存放的就是最后一个子协程对应的函数地址（即输出 9的那个子协程）
		2) 当开始执行子协程对应的函数时，首先执行_p_.runnext 对应的函数，然后按先进先出的顺序执行队列_p_.runq 里的函数。
	*/
	time.Sleep(time.Second * 2)
}

func case2() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			fmt.Print(idx)
		}(i)
		if i%2 != 0 {
			time.Sleep(time.Microsecond)
		}
	}
	// Out: 1032547698
	// 证明了1中的解释,每两个一组，每组都是后面个先执行
	time.Sleep(time.Second * 2)
}
