package cutfile

import (
	"time"

	"github.com/natefinch/lumberjack"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func NewLfsHook(logPath string, maxAge time.Duration) logrus.Hook {
	// 不同等级日志分别配置切割参数
	infoWriter := &lumberjack.Logger{
		Filename: logPath + "/info_%Y-%m-%d.log",
		MaxAge:   int(maxAge),
	}
	errWriter := &lumberjack.Logger{
		Filename: logPath + "/error_%Y-%m-%d.log",
		MaxAge:   int(maxAge),
	}
	// 将不同等级日志写入不同的文件
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: infoWriter,
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  infoWriter,
		logrus.ErrorLevel: errWriter,
		logrus.FatalLevel: errWriter,
		logrus.PanicLevel: errWriter,
	}, &logrus.TextFormatter{})

	return lfsHook
}
