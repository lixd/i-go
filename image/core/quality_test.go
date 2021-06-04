package core

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_quality(t *testing.T) {
	file, err := ioutil.ReadFile("../assets/origin.png")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := Quality(file, 10)
	create, err := os.Create("./compress.png")
	if err != nil {
		t.Fatal(err)
	}
	defer create.Close()
	_, _ = create.Write(bytes)
}

// 4ms
func BenchmarkQuality(b *testing.B) {
	img, err := ioutil.ReadFile("../assets/origin.png")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, _ = Quality(img, 70)
	}
}
