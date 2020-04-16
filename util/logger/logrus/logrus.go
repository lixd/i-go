package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const (
	// 日志存储位置
	LogPath = "./log/zap/logs"
	// 日志文件最大保存时间
	MaxAge = time.Hour * 24 * 90
	// 日志切割时间间隔
	RotationTime = time.Hour
)

func main() {
	InitLogger()
	simpleHttpGet("http://www.baidu.com")
	simpleHttpGet("http://www.google.com")
}

func InitLogger() {
	hook := newLfsHook(LogPath, MaxAge, RotationTime, &logrus.JSONFormatter{})
	// 添加hook
	logrus.AddHook(hook)

	// 设置日志格式为json格式 如果添加了hook 这里的设置可能会被hook中的覆盖
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)
	// 设置日志级别为trace以上
	logrus.SetLevel(logrus.TraceLevel)
	// 打印`打印日志的方法`
	logrus.SetReportCaller(true)
}

// newLfsHook 构建hook 包含日志切割归档功能
func newLfsHook(logPath string, maxAge time.Duration, rotationTime time.Duration, formatter logrus.Formatter) logrus.Hook {
	// 不同等级日志分别配置切割参数
	infoWriter, err := rotatelogs.New(
		logPath+"/info_%Y-%m-%d.log",
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
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scenes": "Error fetching url..", "url": url}).Error(err)
	} else {
		logrus.WithFields(logrus.Fields{"statusCode": resp.Status, "url": url}).Info("Success..")
		resp.Body.Close()
	}
}
