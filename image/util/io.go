package util

import (
	"image"
	"os"

	"github.com/pkg/errors"
)

// LoadImage 从指定位置加载图片
func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "打开文件")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	return img, err
}

// SaveImage 保存图片到指定位置
func SaveImage(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "创建文件")
	}
	defer f.Close()
	bytes, err := Image2Bytes(img)
	if err != nil {
		return errors.Wrap(err, "image对象转二进制数据")
	}
	_, _ = f.Write(bytes)
	_ = f.Sync()
	return err
}
