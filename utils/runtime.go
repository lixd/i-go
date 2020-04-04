package utils

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

// Trace 测试func运行时间
func Trace(msg string) func() {
	start := time.Now()
	logrus.Printf("-------------------------enter %s--------------------------", msg)
	return func() {
		logrus.Printf("--------------------exit %s (%s)--------------------", msg, time.Since(start))
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
