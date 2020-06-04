package izap

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"strings"
	"time"
)

var (
	Logger *zap.SugaredLogger
)

const (
	// LogPath 日志文件路径
	LogPath = "./log/zap/logs"
)

// InitLogger logger对象初始化
/*
包括 encoder writer 日志文件切割和分等级存储等
*/
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
	// 实现两个判断日志等级的interface
	// 如果每个级别的日志都需要分开输出的话 这里在加几个即可
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		//  warn及其以下的都打印在info中
		return lvl <= zapcore.WarnLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		// error及其以上的都打印在error中
		return lvl >= zapcore.ErrorLevel
	})
	// 获取 info、error日志文件的io.Writer
	infoWriter := getWriter(logPath + "/info.log")
	errorWriter := getWriter(logPath + "/error.log")
	encoder := getEncoder()
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		// 分别指定writer和level
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)
	// 传入 zap.AddCaller()显示打日志点的文件名和行数
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// ISO8601 UTC 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ISO8601 UTC 时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
	// 可选普通Encoder
	// return zapcore.NewConsoleEncoder(encoderConfig)
}

// getWriter 传入日志文件存储地址 返回一个writer
func getWriter(logPath string) io.Writer {
	// 生成rotatelogs的Writer
	// 拼接日志文件格式 e.g.:error-2020-03-13-16.log
	fullLogPath := strings.Replace(logPath, ".log", "", -1) + "-%Y-%m-%d-%H.log"
	writer, err := rotatelogs.New(
		fullLogPath,
		rotatelogs.WithLinkName(logPath+"latest.log"), // 生成软链,指向最新日志文件
		// WithMaxAge和WithRotationCount 只能同时指定一个 否则会panic
		rotatelogs.WithMaxAge(time.Hour*24*7), // 日志最大保存时间
		// rotatelogs.WithRotationCount(10),       // 日志文件最大保存数
		rotatelogs.WithRotationTime(time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return writer
}

func GetLogPath() string {
	logPath := viper.GetString("logPath")
	return logPath
}
