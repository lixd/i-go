package main

import (
	"fmt"
	"math/rand"
)

// 正确使用姿势:程序启动时初始化设定一次随机种子即可，比如直接把当前纳秒时间戳当做随机数种子。
func main() {
	// randOne()
	randTwo()
	// randThree()
}

func randOne() {
	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(10)) // 每轮10个随机数都是按照 1779185060 这个顺序出现
	}
}
func randTwo() {
	rand.Seed(10)
	for i := 0; i < 20; i++ {
		fmt.Print(rand.Intn(10)) // 每轮10个随机数都是按照 4879849558 这个顺序出现
	}
}
func randThree() {
	for i := 0; i < 10; i++ {
		rand.Seed(10)
		fmt.Print(rand.Intn(10)) // 每次随机值都为4
	}
}
