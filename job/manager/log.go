package manager

import (
	"i-go/job/model"
	"i-go/job/repository"
	"time"
)

func Log(log *model.JobLog) {
	//if r := recover();r != nil {
	//	log.EndTime = time.Now().Unix()
	//	log.ErrorMessage = fmt.Sprintf("每日结算错误，未知异常,%v",r)
	//	log.Status = JobStatus.Error
	//	_ = repository.Log.InsertJobLog(log)
	//}
	log.EndTime = time.Now().Unix()
	_ = repository.Log.InsertJobLog(log)
}
