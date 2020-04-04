package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobLog struct {
	Id             primitive.ObjectID `bson:"_id"`
	JobId          string             `bson:"JobId"`          //任务id
	JobCode        string             `bson:"JobCode"`        //指定code 用于区别不同任务
	JobName        string             `bson:"JobName"`        //任务名字
	Desc           string             `bson:"Desc"`           //任务描述信息
	StartTime      int64              `bson:"StartTime"`      //任务开始时间
	EndTime        int64              `bson:"EndTime"`        //任务结束时间
	ExecuteMachine string             `bson:"ExecuteMachine"` //执行的机器
	Status         int                `bson:"Status"`         //任务状态
	CronSpec       string             `bson:"CronSpec"`       //执行频率
	ErrorMessage   string             `bson:"ErrorMessage"`   //返回错误信息
}

func (*JobLog) GetCollectionName() string {
	return "JobLog"
}
