package lock

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3/concurrency"
	"i-go/core/etcd"
	"sync"
	"time"
)

var (
	client = etcd.CliV3
)

/*
etcd分布式锁
在etcd事务中查询key的Createrevision是否为0，等于0则创建key和value，表示抢锁成功；
不等于0则for循环watch 一直阻塞到delete事件触发 然后在创建key并返回。
如果需要提前返回只能手动传context来控制超时
*/

// newMutex 获取concurrency.Mutex 每次抢锁都需要重新获取session
func newMutex(ttl int, pfx string) *concurrency.Mutex {
	// NewSession的时候指定TTL为10秒(默认TTL为60秒) 即10秒后锁会自动释放
	session, err := concurrency.NewSession(client, concurrency.WithTTL(ttl))
	if err != nil {
		logrus.WithFields(logrus.Fields{"Scenes": "etcd NewSession"}).Error(err)
	}
	// NewMutex 指定将锁存在哪里 需要强锁的客户端必须指定同一个prefix
	mutex := concurrency.NewMutex(session, pfx)
	return mutex
}

// testLock 测试多协程抢锁
func testLock(ttl int, pfx string, num int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	mutex := newMutex(ttl, pfx)
	err := mutex.Lock(context.Background())
	if err != nil {
		fmt.Printf("协程:%v 抢锁失败\n", num)
	} else {
		fmt.Printf("协程:%v 抢锁成功\n", num)
		doSomething()
		_ = mutex.Unlock(context.Background())
		fmt.Printf("协程:%v 释放锁\n", num)
	}
}

// doSomething 延时模拟业务逻辑
func doSomething() {
	time.Sleep(time.Second)
}
