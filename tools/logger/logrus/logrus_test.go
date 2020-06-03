package zap

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestLogrus(t *testing.T) {
	InitLogger()
	simpleHttpGet("http://www.baidu.com")
	simpleHttpGet("http://www.google.com")
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
