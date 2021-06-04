package core

import (
	"bytes"
	"image"
	"image/jpeg"
)

// Quality 图片质量调整
func Quality(data []byte, q int) ([]byte, error) {
	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: q})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
