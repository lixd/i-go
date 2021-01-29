package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_quality(t *testing.T) {
	img, err := ioutil.ReadFile("./origin.png")
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := Quality(img, 10)
	create, err := os.Create("./compress.png")
	if err != nil {
		fmt.Println(err)
	}
	defer create.Close()

	create.Write(bytes)
}

// 4ms
func BenchmarkQuality(b *testing.B) {
	img, err := ioutil.ReadFile("./origin.png")
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < b.N; i++ {
		_, _ = Quality(img, 70)
	}
}
