package cutfile

import (
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// LogPath 日志存储位置
	LogPath = "./log/zap/logs"
	// MaxAge 日志文件最大保存时间
	MaxAge = time.Hour * 24 * 90
)

func main() {
	hook := NewLfsHook(LogPath, MaxAge)
	// 添加hook
	logrus.AddHook(hook)
	// 会打印出DefaultFieldHook中添加的field  appName=MyAppName
	logrus.Info("") // time="2020-03-13T18:42:08+08:00" level=info appName=MyAppName
}
