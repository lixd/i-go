package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//TestWaitGroup()
	//TestChan()
	//TestWithTimeOut()
	//TestWithCancel()
	//TestWithTimeout()
	TestWithValue()
}

func TestWithTimeOut() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*9)
	res := 0
	for i := 0; i < 10; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			fmt.Println("timeout exit")
			return
		default:
		}
	}
}

func inc(i int) int {
	res := i + 1                // 虽然我只做了一次简单的 +1 的运算,
	time.Sleep(1 * time.Second) // 但是由于我的机器指令集中没有这条指令,
	// 所以在我执行了 1000000000 条机器指令, 续了 1s 之后, 我才终于得到结果。B)
	fmt.Printf("res= %d \n", res)
	return res
}

func TestChan() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("程序退出")
				return
			default:
				fmt.Println("程序运行中。。")
				time.Sleep(time.Second * 2)
			}
		}
	}()
	time.Sleep(time.Second * 10)
	stop <- true
	fmt.Println("停止程序 指令")
	time.Sleep(5 * time.Second)
}

func TestWaitGroup() {
	// 类似 Java 中的 CyclicBarrier ？
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("好了，大家都干完了，放工")
}
func TestWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	add := CountAddCancel(ctx)
	for value := range add {
		if value > 30 {
			cancel()
			break
		}
	}
	fmt.Println("正在统计结果。。。")
	time.Sleep(1500 * time.Millisecond)
}
func TestWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	CountAddTimeOut(ctx)
	defer cancel()
}
func CountAddCancel(ctx context.Context) <-chan int {
	c := make(chan int)
	n := 0
	t := 0
	go func() {
		for {
			time.Sleep(time.Second * 1)
			select {
			case <-ctx.Done():
				fmt.Printf("耗时 %d S 累加值 % d \n", t, n)
				return
			case c <- n:
				// 随机增加1-5
				incr := rand.Intn(4) + 1
				n += incr
				t++
				fmt.Printf("当前累加值 %d \n", n)
			}
		}
	}()
	return c
}
func CountAddTimeOut(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("时间到了 \n")
			return
		default:
			incr := rand.Intn(4) + 1
			n += incr
			fmt.Printf("当前累加值 %d \n", n)
		}
		time.Sleep(time.Second)
	}
}
func TestWithValue() {
	ctx := context.WithValue(context.Background(), "id", "123456")
	id := ctx.Value("id")
	fmt.Printf("id= %v \n", id)
}
