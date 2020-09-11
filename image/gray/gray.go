package main

import (
	"bytes"
	"errors"
	"github.com/Comdex/imgo"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

const (
	path = "D:\\usr\\img\\origin.png"
	save = "D:\\usr\\img\\gray.png"
)

func main() {
	local()
}
func local() {
	// 1.读取数据
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 2.灰度化
	grays := RGB2Gray(file)
	// 3.保存
	create, err := os.Create(save)
	if err != nil {
		panic(err)
	}
	defer create.Close()
	create.Write(grays)
}

// RBGGray 图像灰度处理-基于人眼感知算法
func RGB2Gray(file io.Reader) (gray []byte) {
	// 1.解析图片加载矩阵
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	imgMatrix, err := imgo.Read(img)
	if err != nil {
		panic(err)
	}
	// 2.灰度化
	rgb2Gray := doRGB2Gray(imgMatrix)
	// 3.矩阵转为[]byte
	gray, err = matrix2Bytes(rgb2Gray)
	if err != nil {
		panic(err)
	}
	return
}
func doRGB2Gray(src [][][]uint8) [][][]uint8 {
	height := len(src)
	width := len(src[0])
	imgMatrix := imgo.NewRGBAMatrix(height, width)
	copy(imgMatrix, src)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var (
				red   = imgMatrix[i][j][0]
				green = imgMatrix[i][j][1]
				blue  = imgMatrix[i][j][2]
			)

			gray := uint8(float64(red)*0.3 + float64(green)*0.59 + float64(blue)*0.11)
			imgMatrix[i][j][0] = gray
			imgMatrix[i][j][1] = gray
			imgMatrix[i][j][2] = gray
		}
	}
	return imgMatrix
}

// matrix2Bytes 矩阵转成[]byte
func matrix2Bytes(imgMatrix [][][]uint8) ([]byte, error) {
	height := len(imgMatrix)
	width := len(imgMatrix[0])

	if height == 0 || width == 0 {
		return nil, errors.New("the input of matrix is illegal")
	}

	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nrgba.SetNRGBA(j, i, color.NRGBA{R: imgMatrix[i][j][0], G: imgMatrix[i][j][1], B: imgMatrix[i][j][2], A: imgMatrix[i][j][3]})
		}
	}
	buf := new(bytes.Buffer)
	err := png.Encode(buf, nrgba)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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
	imgMatrix := imgo.NewRGBAMatrix(height, width)
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
func max(nums ...uint8) uint8 {
	var max uint8
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func min(nums ...uint8) uint8 {
	var min uint8
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}
