package core

import (
	"testing"

	"i-go/image/util"
)

func Test_quality(t *testing.T) {
	image, err := util.LoadImage(originImage)
	if err != nil {
		t.Fatal("LoadImage", err)
	}
	quality, err := Quality(image, 10)
	err = util.SaveImage("./quality.jpg", quality)
	if err != nil {
		t.Fatal("SaveImage", err)
	}
}

// 4079862 ns/op 4ms
func BenchmarkQuality(b *testing.B) {
	image, err := util.LoadImage(originImage)
	if err != nil {
		b.Fatal("LoadImage", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Quality(image, 70)
	}
}
