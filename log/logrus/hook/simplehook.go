package main

import "github.com/sirupsen/logrus"

// DefaultFieldHook 需要实现Fire和Levels接口
type DefaultFieldHook struct {
}

// Fire 修改 logrus.Entry 具体的hook代码
func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = "MyAppName"
	return nil
}

// Levels 这里返回触发Fire方法的日志等级
func (hook *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
