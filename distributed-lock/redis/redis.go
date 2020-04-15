package main

import (
	_ "i-go/core/conf"
	_ "i-go/core/db/redisdb"

	"github.com/sirupsen/logrus"
	"i-go/util/lock"
	"i-go/utils"
	"strconv"
	"sync"

	"time"
)

const (
	MyLock = "mylock"
	Expire = time.Second
)

var waitGroup sync.WaitGroup

func main() {
	for i := 1; i <= 10; i++ {
		waitGroup.Add(1)
		go checkLock(strconv.Itoa(i))
	}
	waitGroup.Wait()
	logrus.Info("所有任务执行完成...")
}

func checkLock(taskId string) {
	uuid := utils.StringHelper.GetUUID()
	for {
		if lock.RedisLock.GetLock(MyLock, uuid, Expire) {
			logrus.Infof("任务%s 获取锁成功", taskId)
			doSomething(taskId)
			lock.RedisLock.ReleaseLock(MyLock, uuid)
			logrus.Infof("任务%s 释放锁", taskId)
			break
		} else {
			logrus.Infof("任务%s 获取锁失败", taskId)
		}
		time.Sleep(time.Millisecond * 300)
	}
	waitGroup.Done()
}

func doSomething(taskId string) {
	logrus.Infof("%s 任务执行中...", taskId)
	// 模拟业务逻辑延时
	time.Sleep(time.Second * 2)
}
