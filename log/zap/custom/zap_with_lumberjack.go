package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	customLogger *zap.SugaredLogger
)

// LogPath 日志文件路径
const (
	LogPathInfo  = "./log/zap/logs/info.log"
	LogPathError = "./log/zap/logs/error.log"
)

// zap日志框架+lumberjack做文件切割
// 只能按照文件大小切割 不是很舒服 能按时间切割就好了
func main() {
	InitLogger()
	defer customLogger.Sync()
	// simpleHttpGet("www.google.com")
	for i := 0; i < 9999; i++ {
		simpleHttpGet("http://www.baidu.com")
	}
}

func InitLogger() {
	// 实现两个判断日志等级的interface
	// 如果每个级别的日志都需要分开输出的话 这里再加几个即可
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		// info中只打印 info 和 warn
		return lvl <= zapcore.WarnLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		// error 及其以上的都打印在 error 中
		return lvl >= zapcore.ErrorLevel
	})

	infoWriter := getLogWriter(LogPathInfo)
	errorWriter := getLogWriter(LogPathError)
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)
	// 最后创建具体的Logger
	// 传入 zap.AddCaller()显示打日志点的文件名和行数
	logger := zap.New(core, zap.AddCaller())
	customLogger = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	// 获取默认配置然后自定义调整
	encoderConfig := zap.NewProductionEncoderConfig()
	// ISO8601 UTC 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
	// 普通Encoder
	// return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 创建一个WriterSyncer
// path 文件路径
func getLogWriter(path string) zapcore.WriteSyncer {
	// 带有日志切割功能
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,  // 日志文件的位置
		MaxSize:    1,     // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,     // 保留旧文件的最大个数
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getLogWriterSimple 创建一个WriterSyncer
func getLogWriterSimple(path string) zapcore.WriteSyncer {
	file, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintf("log init getLogWriter error: %v", err))
	}
	return zapcore.AddSync(file)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		customLogger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		customLogger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
