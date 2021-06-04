package utils

import (
	"fmt"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

// Trace 测试func运行时间
func Trace(msg string) func() {
	start := time.Now()
	fmt.Printf("-------------------------enter %s--------------------------\n", msg)
	return func() {
		fmt.Printf("--------------------exit %s (%s)--------------------\n", msg, time.Since(start))
	}
}

// Caller 返回调用函数的名字
func Caller() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// InitLog 记录启动日志
func InitLog(msg string) func() {
	start := time.Now()
	logrus.Infof("-------------------- %s init begin -------------------------------", msg)
	return func() {
		logrus.Infof("-------------------- %s init end   time consuming %v ------", msg, time.Since(start))
	}
}
