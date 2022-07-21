package main

import (
	"strconv"
	"sync"
	"time"

	"i-go/core/conf"
	_ "i-go/core/conf"
	"i-go/core/db/redisdb"
	_ "i-go/core/db/redisdb"

	"i-go/tools/lock"
	"i-go/utils"

	"github.com/sirupsen/logrus"
)

const (
	MyLock = "mylock"
	Expire = time.Second
)

func Init() {
	if err := conf.Load("D:/lillusory/projects/i-go/conf/config.yml"); err != nil {
		panic(err)
	}
	redisdb.Init()
}

func main() {
	Init()

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go checkLock(strconv.Itoa(i), &wg)
	}
	wg.Wait()
	logrus.Info("所有任务执行完成...")
}

func checkLock(taskId string, wg *sync.WaitGroup) {
	defer wg.Done()
	redisLock := lock.NewRedisLock(redisdb.Cli)
	uuid := utils.StringHelper.GetUUID()
	for {
		if redisLock.Lock(MyLock, uuid, Expire) {
			logrus.Infof("任务%s 获取锁成功", taskId)
			doSomething(taskId)
			_ = redisLock.UnLock(MyLock, uuid)
			logrus.Infof("任务%s 释放锁", taskId)
			break
		} else {
			logrus.Infof("任务%s 获取锁失败", taskId)
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func doSomething(taskId string) {
	logrus.Infof("%s 任务执行中...", taskId)
	// 模拟业务逻辑延时
	time.Sleep(time.Second * 2)
}
