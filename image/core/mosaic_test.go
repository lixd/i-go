package core

import (
	"testing"

	"i-go/image/util"
)

func TestOverall(t *testing.T) {
	image, err := util.LoadImage(originImage)
	if err != nil {
		t.Fatal("LoadImage", err)
	}
	mosaic := Overall(image, 3, 3)
	if err != nil {
		t.Fatal("Overall", err)
	}
	err = util.SaveImage("./mosaic-overall.jpg", mosaic)
	if err != nil {
		t.Fatal("SaveImage", err)
	}
}

func TestPartial(t *testing.T) {
	image, err := util.LoadImage(originImage)
	if err != nil {
		t.Fatal("LoadImage", err)
	}
	mosaic := Partial(image, 300, 200)
	if err != nil {
		t.Fatal("Overall", err)
	}
	err = util.SaveImage("./mosaic-partial.jpg", mosaic)
	if err != nil {
		t.Fatal("SaveImage", err)
	}
}

// 1889303 ns/op 18ms
func BenchmarkOverall(b *testing.B) {
	image, err := util.LoadImage(originImage)
	if err != nil {
		b.Fatal("LoadImage", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Overall(image, 3, 3)
	}
}

// 3546414 ns/op 3ms
func BenchmarkPartial(b *testing.B) {
	image, err := util.LoadImage(originImage)
	if err != nil {
		b.Fatal("LoadImage", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Partial(image, 100, 60)
	}
}
