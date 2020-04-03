package manager

import (
	"github.com/sirupsen/logrus"
	"i-go/job/model"
	"i-go/job/repository"
)

// CreateJob
func CreateJob(taskCode, name, spec, desc string, duration int64) model.JobInfo {
	info, err := repository.Job.Init(taskCode, name, spec, desc, duration)
	if err != nil {
		logrus.Errorf("init service fail", name)
		panic(err)
	}
	return info
}
