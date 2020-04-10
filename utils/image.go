package utils

import (
	"github.com/nfnt/resize"
	"image"
)

// ImageResize 缩放 w,h 为缩放后图片宽高
func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
