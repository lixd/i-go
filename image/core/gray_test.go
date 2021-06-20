package core

import (
	"os"
	"testing"
)

func TestRGB2Gray(t *testing.T) {
	// 1.读取数据
	file, err := os.Open(originImage)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 2.灰度化
	grays := RGB2Gray(file)
	// 3.保存
	create, err := os.Create("./gray.jpg")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	_, _ = create.Write(grays)
}
