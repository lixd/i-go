package core

import (
	"i-go/job/model"
	"time"
)

type IJob interface {
	Init()
	Work()
	WorkDirectly(t time.Time)
	GetInfo() model.JobInfo
	CheckAndUpdate() bool
}
