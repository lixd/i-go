package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var done = false

func main() {
	// demo1()
	demo2()
	// demo3()
}

/*
for i in {1..100}; do
curl https: // www.puug.com/api/v1/funding/list?page=1&sort=+&cateId=
done
*/

/*
查询缓存：
1.列表页 首页、众筹中、求解中
2.问题详情
3.猜你喜欢
4.个人页 动态、我的发布、我的回答、收藏、打赏
移除缓存：
题主：更新问题 影响问题内容、价格、状态
题主：追加赏金 影响问题赏金
其他人：参与众筹、追加众筹、撤资 影响问题阶段、众筹进度、
其他人：收藏问题 影响问题收藏数--可以缓存收藏数，cronJob定时增加
其他人：打赏 影响问题分享数（可能，如果带有referer字段）--可以缓存分享数，cronJob定时增加
打赏人：评价(赞或踩) 会影响收益相关字段
admin：置顶、恢复/隐藏、关闭、删除
*/
func demo1() {
	var (
		locker sync.Mutex
		cond   = sync.NewCond(&locker)
		wg     sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(number int) {
			// wait()方法内部是先释放锁 然后在加锁 所以这里需要先 Lock()
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait() // 等待通知,阻塞当前 goroutine
			fmt.Printf("g %v ok~ \n", number)
			wg.Done()
		}(i)
	}

	for i := 0; i < 5; i++ {
		// 每过 50毫秒 唤醒一个 goroutine
		cond.Signal()
		time.Sleep(time.Millisecond * 50)
	}

	time.Sleep(time.Millisecond * 50)
	// 剩下5个 goroutine 一起唤醒
	cond.Broadcast()
	fmt.Println("Broadcast...")
	wg.Wait()
}

func demo2() {
	var (
		locker sync.Mutex
		cond   = sync.NewCond(&locker)
	)

	go read1("reader1", cond)
	go read1("reader2", cond)
	go read1("reader3", cond)
	writeSignal("writer1", cond)
	writeSignal("writer2", cond)
	writeSignal("writer3", cond)

	time.Sleep(time.Second)
}

func demo3() {
	cond := sync.NewCond(&sync.Mutex{})

	go read2("reader1", cond)
	go read2("reader2", cond)
	go read2("reader3", cond)
	writeBroadcast("writer", cond)

	time.Sleep(time.Second)
}

func read1(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		fmt.Println(name, "wait")
		// NOTE: 这是因为当 Boardcast 唤醒时，有可能其他 goroutine 先于当前 goroutine 唤醒并抢到锁，导致轮到当前 goroutine 抢到锁的时候，可能条件又不再满足了。
		// 因此，需要在 Wait 返回之后再判断一次是否满足条件，最简单的就是直接将条件检查放在 for 循环中。
		// 因为虽然wait之前调用了Lock 但是Wait方法中会调用 Unlock，这中间可能导致done变量被修改 比如在 Read 之后可以把 done 又切换回 false
		c.Wait()
	}
	log.Println(name, "starts reading")
	done = false
	c.L.Unlock()
}

func writeBroadcast(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func read2(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		fmt.Println(name, "wait")
		// NOTE: 这是因为当 Boardcast 唤醒时，有可能其他 goroutine 先于当前 goroutine 唤醒并抢到锁，导致轮到当前 goroutine 抢到锁的时候，可能条件又不再满足了。
		// 因此，需要在 Wait 返回之后再判断一次是否满足条件，最简单的就是直接将条件检查放在 for 循环中。
		// 因为虽然wait之前调用了Lock 但是Wait方法中会调用 Unlock，这中间可能导致done变量被修改 比如在 Read 之后可以把 done 又切换回 false
		c.Wait()
	}
	log.Println(name, "starts reading")
	// done = false // 唤醒全部时这里就不修改资源状态了 和 read1 区别开
	c.L.Unlock()
}
func writeSignal(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	c.L.Lock()
	done = true
	c.L.Unlock()
	c.Signal()
	log.Println(name, "wakes one")
}
