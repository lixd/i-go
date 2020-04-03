package manager

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"i-go/db/mongodb"
	jobstatus "i-go/job/constant"
	"i-go/job/core"
	"i-go/job/model"
	"math"
	"time"
)

/*
 一个简单的分布式定时任务框架
 主要借助数据库锁实现定时任务的多机器运行
(先修改数据库再运行任务 没有修改成功的机器则不会执行)
*/

// JobManager
type JobManager struct {
	Cron *cron.Cron
	Jobs []core.IJob
}

// DefaultJobManager
var DefaultJobManager = &JobManager{}

// CreateJobs
func (c *JobManager) CreateJobs(jobs []core.IJob) {
	for _, j := range jobs {
		j.Init()
		info := j.GetInfo()
		spec := info.CronSpec
		_ = c.Cron.AddFunc(spec, j.Work)
		c.Jobs = append(c.Jobs, j)
	}
}

// StartAll
func (c *JobManager) StartAll() {
	c.Cron.Start()
}

// Execute
func (c *JobManager) Execute(name string) {
	for _, _job := range c.Jobs {
		if _job.GetInfo().JobName == name {
			_job.WorkDirectly(time.Now())
		}
	}
}

// Done 任务运行完之后修改数据库到 `待执行` 状态 同时执行时间和次数等信息
func (c *JobManager) Done(jobCode string, start, end int64) {
	var jobInfo model.JobInfo

	var filter = bson.D{
		{"JobCode", jobCode},
		{"Status", jobstatus.Process},
	}

	var update = bson.D{
		{"$set", bson.D{
			{"Status", jobstatus.Wait},
			{"LastExecuteDuration", end - start}, // 执行定时任务所花费的时间
		}},
		{"$inc", bson.D{
			{"TotalExecuted", 1}, // 执行次数加1
		}},
	}
	_, _ = mongodb.GetJobCollection(&jobInfo).UpdateOne(nil, filter, update)
}

// CheckAndUpdate 检测当前任务是否需要运行 如果需要会更新数据库中的任务状态
func (c *JobManager) CheckAndUpdate(jobCode, ip string, duration int64) bool {
	var now = time.Now().Unix()
	// 1. 查看当前任务是否处于待执行状态 是则更新到运行中 且在3处返回true 开始执行任务
	var filter = bson.D{
		{"JobCode", jobCode},
		{"Status", jobstatus.Wait},
	}

	var update = bson.D{
		{"$set", bson.D{
			{"Status", jobstatus.Process},
		}},
	}
	var job model.JobInfo
	var opts = options.FindOneAndUpdate().SetReturnDocument(options.Before) // 返回修改之前的文档
	var coll = mongodb.GetJobCollection(&job)

	err := coll.FindOneAndUpdate(nil, filter, update, opts).Decode(&job)
	// 当前任务不在等待阶段， 任务可能已经开始 或者上次任务出现错误，未能正常结束
	if err != nil {
		// 2.1 通过JobCode查询该任务信息
		var filter = bson.M{"JobCode": jobCode}
		err := coll.FindOne(nil, filter).Decode(&job)
		if err != nil {
			return false // 没有找到该任务
		}
		var d = now - job.LastExecuteTime // 距上次执行时间
		logrus.Info("上次执行在", d, "秒前", math.Max(0, float64(job.MinDuration-d)), "秒后可以执行")

		// 程序没有启动过 或者 上次任务执行时间超过最小间隔时间 可能卡死了
		if job.LastExecuteTime == 0 || d >= job.MinDuration {
			// 重新启动任务 且返回true 开始执行任务
			var update = bson.D{
				{"$set", bson.D{
					{"Status", jobstatus.Process},
					{"LastExecuteIP", ip},
					{"LastExecuteTime", now},
				}},
			}
			if _, err := coll.UpdateOne(nil, filter, update); err != nil {
				return false // 重新启动任务出错
			}
			return true // 程序卡死，需要重新执行
		} else {
			//具体什么时候运行由cron表达式决定 这里主要是防止多机器时重复执行的问题
			return false // 程序正常运作，当前不需要执行
		}
	}

	//3. 如果1中没出现问题说明当前任务处于等待状态 可以执行
	// 更新任务信息 然后返回true 执行任务
	filter = bson.D{
		{"_id", job.Id},
	}
	// 执行 记录执行时间和执行ip
	update = bson.D{
		{"$set", bson.D{
			{"LastExecuteIP", ip},
			{"LastExecuteTime", now},
		}},
	}
	_, _ = coll.UpdateOne(nil, filter, update)

	return true // 任务在等待阶段 需要执行

}

func Init() {
	logrus.Info("init start")
	DefaultJobManager.Cron = cron.New()
	logrus.Info("init end")
}
