package main

import "github.com/sirupsen/logrus"

func main() {
	// 添加hook
	logrus.AddHook(new(DefaultFieldHook))
	// 会打印出DefaultFieldHook中添加的field  appName=MyAppName
	logrus.Info("") // time="2020-03-13T18:42:08+08:00" level=info appName=MyAppName
}
