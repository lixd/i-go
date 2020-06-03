package zap

import (
	"go.uber.org/zap"
	"net/http"
	"testing"
)

func TestLogrus(t *testing.T) {
	InitLogger()
	defer LoggerZ.Sync()
	simpleHttpGet("http://www.baidu.com")
	simpleHttpGet("http://www.google.com")
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		LoggerZ.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		LoggerZ.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
