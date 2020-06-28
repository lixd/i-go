package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//TestChan()
	//TestWithCancel()
	//TestWithTimeout()
	//TestWithValue()
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
	ctxName := context.WithValue(ctx, "name", "17x")
	id := ctx.Value("id")
	// ctxName 中没有 id 值 所以会往父节点查找 取 ctx 中的 id 值
	id2 := ctxName.Value("id")
	fmt.Printf("id= %v id2= %v \n", id, id2)
}
