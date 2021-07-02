package core

import (
	"image"

	"github.com/pkg/errors"
	"i-go/image/util"
)

// RGB2Gray 图像灰度处理-基于人眼感知算法
func RGB2Gray(origin image.Image) (image.Image, error) {
	// 1.解析图片加载矩阵
	// imgMatrix, err := imgo.Read(origin)
	imgMatrix := util.Image2Matrix(origin)
	// 2.灰度化
	rgb2Gray := doRGB2Gray(imgMatrix)
	// 3.矩阵转为image
	gray, err := util.Matrix2Image(rgb2Gray)
	return gray, errors.Wrap(err, "矩阵转image")
}

// doRGB2Gray
func doRGB2Gray(src [][][]uint8) [][][]uint8 {
	height := len(src)
	width := len(src[0])
	imgMatrix := util.NewRGBAMatrix(height, width)
	// imgMatrix := imgo.NewRGBAMatrix(height, width)
	copy(imgMatrix, src)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var (
				red   = imgMatrix[i][j][0]
				green = imgMatrix[i][j][1]
				blue  = imgMatrix[i][j][2]
			)
			// 基于人眼感知算法计算灰度值
			gray := uint8(float64(red)*0.3 + float64(green)*0.59 + float64(blue)*0.11)
			// RGB3通道值一致时 图片就是灰色
			imgMatrix[i][j][0] = gray
			imgMatrix[i][j][1] = gray
			imgMatrix[i][j][2] = gray
		}
	}
	return imgMatrix
}

// RGB2GrayDemo 灰度化具体算法
/*
具体逻辑 根据RBG三个值计算出一个gray值 然后替换RGB三个值
参考链接
https://segmentfault.com/a/1190000009000216
https://tannerhelland.com/2011/10/01/grayscale-image-algorithm-vb6.html
取 gray值得算法有多种 常见的如下
1.平均法
2.基于人眼感知（公式）
3.去饱和度算法
4.分解法
5.单通道
6.自定义灰度阴影
*/
func RGB2GrayDemo(src [][][]uint8) [][][]uint8 {
	height := len(src)
	width := len(src[0])
	// imgMatrix := imgo.NewRGBAMatrix(height, width)
	imgMatrix := util.NewRGBAMatrix(height, width)

	copy(imgMatrix, src)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// avg := (imgMatrix[i][j][0] + imgMatrix[i][j][1] + imgMatrix[i][j][3]) / 3
			var (
				red   = imgMatrix[i][j][0]
				green = imgMatrix[i][j][1]
				blue  = imgMatrix[i][j][2]
			)
			// 1.平均法
			// gray := (red + green + blue) / 3
			// 2.基于人眼感知
			// 2.1
			gray := uint8(float64(red)*0.3 + float64(green)*0.59 + float64(blue)*0.11)
			// 2.2
			// gray := uint8(float64(red)*0.2126 + float64(green)*0.7152 + float64(blue)*0.0722)
			// 2.3
			// gray := uint8(float64(red)*0.299 + float64(green)*0.587 + float64(blue)*0.114)

			// 3.去饱和度算法
			// gray := (max(red, green, blue) + min(red, green, blue)) / 2
			// 4.分解法
			// 4.1 最大值分解
			// gray := max(red, green, blue)
			// 4.2 最小值分解
			// gray := min(red, green, blue)
			// 5.单通道
			// gray := red
			// gray := green
			// gray := blue
			imgMatrix[i][j][0] = gray
			imgMatrix[i][j][1] = gray
			imgMatrix[i][j][2] = gray
		}
	}
	return imgMatrix
}
