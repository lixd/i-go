package logger

import (
	"net/http"
	"testing"

	"go.uber.org/zap"
)

func TestInitLogger(t *testing.T) {
	InitLogger()
	simpleHttpGet("https://www.vaptcha.com")
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		ILog.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		ILog.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		_ = resp.Body.Close()
	}
}
