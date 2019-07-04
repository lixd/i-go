package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

/*
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
*/

// 定时任务 github.com/robfig/cron
func main() {
	c := cron.New()
	job := job{"1", "cron job"}
	// 直接添加Func
	err := c.AddFunc("1/2 * * * * *", func() { fmt.Println("cron") })
	err = c.AddFunc("@every 0h0m1s", func() { fmt.Println("cron special") })
	// 添加一个Job对象 是一个接口 只需实现Run方法
	err = c.AddJob("@every 0h0m2s", &job)
	if err != nil {
		fmt.Printf("AddFunc err =%v \n", err)
	}
	// c.Start()
	for i, v := range c.Entries() {
		fmt.Printf("index=%v  \n", i)
		fmt.Printf("value.Job=%v  \n", v.Job)           // job
		fmt.Printf("value.next=%v  \n", v.Next)         // 下次执行时间
		fmt.Printf("value.Prev=%v  \n", v.Prev)         // 上次执行时间
		fmt.Printf("value.Schedule=%v  \n", v.Schedule) // 调度 执行间隔
	}
	time.Sleep(time.Second * 40)
}

type job struct {
	id   string
	desc string
}

// 实现 Run 方法 即实现了Job 接口
func (job *job) Run() {
	fmt.Printf("job id=%v desc=%v \n", job.id, job.desc)
}
