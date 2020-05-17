package lock

import (
	"sync"
	"testing"
)

/*
协程 1 抢锁成功
协程 1 释放锁
协程 0 抢锁成功
协程 0 释放锁
协程 2 抢锁成功
协程 2 释放锁

正常
*/
func TestMyEtcdMutex_Lock(t *testing.T) {
	var waitGroup = new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go testLockSync(10, "/mylock", i, waitGroup)
	}
	waitGroup.Wait()
}
