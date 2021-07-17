package lock

import (
	"sync"
	"testing"
)

/*
协程:2 抢锁成功
协程:2 释放锁
协程:1 抢锁成功
协程:1 释放锁
协程:0 抢锁成功
协程:0 释放锁
正常
*/
func TestEtcdMutex_Lock(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go testLock(10, "/mylock", i, &wg)
	}
	wg.Wait()
}
