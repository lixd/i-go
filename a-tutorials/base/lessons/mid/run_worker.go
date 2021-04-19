package main

import (
	"log"
	"os"
	"time"

	common2 "i-go/a-tutorials/base/lessons/mid/common"
)

func main() {
	log.Println("...开始执行任务...")

	timeout := 3 * time.Second
	r := common2.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case common2.ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case common2.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
