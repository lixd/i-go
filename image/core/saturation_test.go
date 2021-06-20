package core

import (
	"os"
	"testing"
)

const (
	originImage = "../assets/origin.png"
)

func TestAdjustSaturation(t *testing.T) {
	file, err := os.Open(originImage)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	// 2.灰度化
	grays := AdjustSaturation(file, 0.6)
	// 3.保存
	create, err := os.Create("./adjustSaturation.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer create.Close()
	_, _ = create.Write(grays)
}
