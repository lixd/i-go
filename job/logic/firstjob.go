package logic

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	jobstatus "i-go/job/constant"
	"i-go/job/manager"
	"i-go/job/model"
	"i-go/job/service/secondly"
	"i-go/utils"
	"time"
)

/*
	第一个定时任务
*/
// FirstJob实现IJob接口
type FirstJob struct {
	JobInfo model.JobInfo
}

func (c *FirstJob) Init() {
	c.JobInfo = manager.CreateJob("2020040301", "第一个定时任务", "0/5 * * * * ?", "每5秒执行一次", 60)
}

func (c *FirstJob) Work() {
	if c.CheckAndUpdate() {
		c.work(time.Now())
	}
}

func (c *FirstJob) WorkDirectly(t time.Time) {
	// 这里也需要执行CheckAndUpdate() 用于同步更新任务状态
	c.CheckAndUpdate()
	c.work(t)
}

func (c *FirstJob) GetInfo() model.JobInfo {
	return c.JobInfo
}

func (c *FirstJob) CheckAndUpdate() bool {
	return manager.DefaultJobManager.CheckAndUpdate(c.JobInfo.JobCode, utils.GetIntranetIP(), c.JobInfo.MinDuration)
}

func (c *FirstJob) work(t time.Time) {
	var start = time.Now().UnixNano() / 1e6
	var log = &model.JobLog{
		JobId:          c.JobInfo.Id.Hex(),
		JobCode:        c.JobInfo.JobCode,
		JobName:        c.JobInfo.JobName,
		CronSpec:       c.JobInfo.CronSpec,
		Desc:           c.JobInfo.Desc,
		StartTime:      time.Now().Unix(),
		ExecuteMachine: utils.GetIntranetIP(),
		Id:             primitive.NewObjectID(),
		Status:         jobstatus.Done,
	}
	defer manager.Log(log)

	secondly.FirstJob()
	end := time.Now().UnixNano() / 1e6
	log.EndTime = end
	manager.DefaultJobManager.Done(c.JobInfo.JobCode, start, end)
}
