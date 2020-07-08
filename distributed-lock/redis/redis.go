package main

import (
	"github.com/go-redis/redis"
	"i-go/core/conf"
	_ "i-go/core/conf"
	"i-go/core/db/redisdb"
	_ "i-go/core/db/redisdb"

	"github.com/sirupsen/logrus"
	"i-go/tools/lock"
	"i-go/utils"
	"strconv"
	"sync"

	"time"
)

const (
	MyLock = "mylock"
	Expire = time.Second
)

var (
	waitGroup sync.WaitGroup
)

func Init() {
	conf.Init("D:/lillusory/projects/i-go/conf/config.yml")
	redisdb.Init()
}

func main() {
	Init()

	for i := 1; i <= 10; i++ {
		waitGroup.Add(1)
		go checkLock(strconv.Itoa(i))
	}
	waitGroup.Wait()
	logrus.Info("所有任务执行完成...")
}

func checkLock(taskId string) {
	redisLock := lock.NewRedisLock(redisdb.RedisClient)
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

// GetUser 获取用户信息
func GetUser(userId int) {
	key := "xxx"
	// 1. 直接查缓存
	m, err := redisdb.RedisClient.HGetAll(key).Result()
	if err != nil && err != redis.Nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller(), "scenes": "获取用户信息，查询缓存"}).Error(err)
	}
	if err == redis.Nil || len(m) == 0 {
		// 2. 如果不存在就去查 数据库
		redisLock := lock.NewRedisLock(redisdb.RedisClient)
		value := time.Now().UnixNano()
		// 2.1 查询数据库需要先获取锁
		isLock := redisLock.Lock(key, value, time.Second*5)
		if isLock {
			// 2.2 成功则查询
			// 2.2.1 查询数据库
			// 2.2.2 同步数据到缓存
		} else {
			// 2.3 失败则 延时后继续重试（第二次执行时可能其他客户端已经把数据同步到缓存了）
			time.Sleep(time.Millisecond * 10)
			GetUser(userId)
		}

	}
	redisdb.RedisClient.Expire(key, time.Minute*30)
}
