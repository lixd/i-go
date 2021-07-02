package core

import (
	"testing"

	"i-go/image/util"
)

func TestRGB2Gray(t *testing.T) {
	// 1.读取数据
	img, err := util.LoadImage(originImage)
	if err != nil {
		t.Fatal("LoadImage err:", err)
	}
	// 2.灰度化
	gray, err := RGB2Gray(img)
	if err != nil {
		t.Fatal("RGB2Gray err:", err)
	}
	// 3.保存
	err = util.SaveImage("./gray.jpg", gray)
	if err != nil {
		t.Fatal("SaveImage err:", err)
	}
}

// 8830345 ns/op 8ms
func BenchmarkGray(b *testing.B) {
	img, err := util.LoadImage(originImage)
	if err != nil {
		b.Fatal("LoadImage err:", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = RGB2Gray(img)
	}
}
