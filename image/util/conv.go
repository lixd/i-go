package util

import (
	"bytes"
	"image"
	"image/jpeg"
)

// Bytes2Image 转换耗时和图片尺寸有关
func Bytes2Image(origin []byte) (image.Image, error) {
	reader := bytes.NewReader(origin)
	return jpeg.Decode(reader)
}

// Image2Bytes 图片转 byte 数组
func Image2Bytes(img image.Image) ([]byte, error) {
	buf := bytes.Buffer{}
	err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
	return buf.Bytes(), err
}
