package zap

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// 日志存储位置
	LogPath = "./log/logrus/logs"
	// 日志文件最大保存时间
	MaxAge = time.Hour * 24 * 90
	// 日志切割时间间隔
	RotationTime = time.Hour * 24
)

func InitLogger(path ...string) {
	var logPath string
	if len(logPath) == 0 {
		logPath = GetLogPath()
	} else {
		logPath = path[0]
	}
	if logPath == "" {
		logPath = LogPath
	}
	hook := newLfsHook(logPath, MaxAge, RotationTime, &logrus.JSONFormatter{})
	// 添加hook
	logrus.AddHook(hook)

	// 设置日志格式为json格式 如果添加了hook 这里的设置可能会被hook中的覆盖
	logrus.SetFormatter(&logrus.JSONFormatter{})
	out, _ := os.OpenFile("./z.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(out)
	// logrus.SetOutput(os.Stdout)
	// 设置日志级别为trace以上
	logrus.SetLevel(logrus.TraceLevel)
	// // 打印`打印日志的方法`
	// logrus.SetReportCaller(true)
}

// newLfsHook 构建hook 包含日志切割归档功能
func newLfsHook(logPath string, maxAge time.Duration, rotationTime time.Duration, formatter logrus.Formatter) logrus.Hook {
	fmt.Println("logPath: ", logPath)
	// 不同等级日志分别配置切割参数
	infoWriter, err := rotatelogs.New(
		logPath+"/info_%Y-%m-%d.log",              // %Y-%m-%d-%H 可以到小时
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	errWriter, err := rotatelogs.New(
		logPath+"/error_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	// 将不同等级日志写入不同的文件
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: infoWriter,
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  infoWriter,
		logrus.ErrorLevel: errWriter,
		logrus.FatalLevel: errWriter,
		logrus.PanicLevel: errWriter,
	}, formatter)

	return lfsHook
}
func GetLogPath() string {
	logPath := viper.GetString("logPath")
	return logPath
}
