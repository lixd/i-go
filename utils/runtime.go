package utils

import (
	"github.com/sirupsen/logrus"
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
