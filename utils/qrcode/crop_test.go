package main

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"testing"
)

func TestClaCoordinate(t *testing.T) {
	//0. 加载图片
	file2, err := os.Open(QrCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "加载图片"}).Error(err)
	}
	all, err := ioutil.ReadAll(file2)
	// 1. 处理
	code := cropCode(all, ImgTypeWeChatMiniProgram)
	// 2. 保存
	f, err := os.OpenFile("D:/wlinno/qrcode_new2.jpg", os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	f.Write(code)
}

// 171毫秒
func BenchmarkClaCoordinate(b *testing.B) {
	file2, err := os.Open(QrCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "加载图片"}).Error(err)
	}
	all, err := ioutil.ReadAll(file2)
	for i := 0; i < b.N; i++ {
		_ = cropCode(all, ImgTypeWeChat)
	}
}
