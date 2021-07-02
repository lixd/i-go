package core

import (
	"testing"

	"i-go/image/util"
)

func TestAdjustSaturation(t *testing.T) {
	image, err := util.LoadImage(originImage)
	if err != nil {
		t.Fatal("LoadImage", err)
	}
	quality, err := AdjustSaturation(image, 0.8)
	err = util.SaveImage("./saturation.jpg", quality)
	if err != nil {
		t.Fatal("SaveImage", err)
	}
}

// 10123722 ns/op 10ms
func BenchmarkSaturation(b *testing.B) {
	image, err := util.LoadImage(originImage)
	if err != nil {
		b.Fatal("LoadImage", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = AdjustSaturation(image, 0.8)
	}
}
