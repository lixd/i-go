package main

import (
	"github.com/sirupsen/logrus"
	"i-go/core/db/mongodb"
	"i-go/core/db/redisdb"
	"i-go/job/core"
	"i-go/job/logic"
	"i-go/job/manager"
	"i-go/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var ts []core.IJob
	ts = append(ts, new(logic.FirstJob))

	manager.Init()
	manager.DefaultJobManager.CreateJobs(ts)
	manager.DefaultJobManager.StartAll()

	logrus.Info("ip:", utils.GetIntranetIp())

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-s:
			logrus.WithFields(logrus.Fields{"main error": "job  shutdown"})
			redisdb.Release()
			mongodb.Release()
		}
	}()

	select {}
}
