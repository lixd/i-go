package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobInfo struct {
	Id                  primitive.ObjectID `bson:"_id"`
	JobCode             string             `bson:"JobCode"`             // 指定code 用于区别不同任务
	JobName             string             `bson:"JobName"`             // 名称
	CreateTime          int64              `bson:"CreateTime"`          // 创建时间
	LastExecuteTime     int64              `bson:"LastExecuteTime"`     // 上次执行时间
	LastExecuteIP       string             `bson:"LastExecuteIP"`       // 上次执行服务器ip
	LastExecuteDuration int64              `bson:"LastExecuteDuration"` // 上次执行花费时间
	Desc                string             `bson:"Desc"`                // 详细描述
	CronSpec            string             `bson:"CronSpec"`            // corn命令
	MinDuration         int64              `bson:"MinDuration"`         // 超时时间
	Status              int                `bson:"Status"`              // 状态
	TotalExecuted       int32              `bson:"TotalExecuted"`       // 执行次数
}

func (*JobInfo) GetCollectionName() string {
	return "JobInfo"
}
