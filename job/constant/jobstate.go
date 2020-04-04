package jobstatus

const (
	//任务信息状态
	Wait    = 1 // 待运行
	Process = 2 // 运行中
	//日志信息状态
	Error //3 执行异常
	Done  //4 执行成功
)
