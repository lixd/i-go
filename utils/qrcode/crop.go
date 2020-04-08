package main

import (
	"bytes"
	"fmt"
	"github.com/Comdex/imgo"
	"github.com/sirupsen/logrus"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 二维码裁剪工具包 根据截图把二维码裁剪出来
// 笨办法 根据像素点查找的 需要清晰一点的
const (
	QrCode     = "res/qrcode_wechat.jpg"
	QrCodeSave = "res/qrcode_crop.jpg"
	// 图片类型 小程序1 微信2
	ImgTypeWeChatMiniProgram = 1
	ImgTypeWeChat            = 2
)

func main() {
	imgBytes := load2Bytes()
	imgCropBytes := cropCode(imgBytes, ImgTypeWeChat)
	saveFile(imgCropBytes, QrCodeSave)
}

func load2Bytes() []byte {
	file2, err := os.Open(QrCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "加载图片"}).Error(err)
	}
	all, err := ioutil.ReadAll(file2)
	return all
}

func saveFile(code []byte, path string) {
	f, err := os.OpenFile(path, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return
	}
	f.Write(code)
}

// cropCode 图片裁剪
//Benchmark 171541200 ns/op
func cropCode(data []byte, imgType int) []byte {
	// 0.加载图片
	origin, err := loadImageFromBytes(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "解析图片"}).Error(err)
		return nil
	}
	// 1. 定位小程序码位置
	w, h, dx := positionDetectionNew(imgType)
	// 2.计算裁剪坐标
	x0, y0, l := ClaCoordinate(w, h, dx, imgType)
	// 3. 裁剪
	imageCopy, err := clip(origin, x0, y0, l, l)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "裁剪图片"}).Error(err)
	}
	// 4. return []byte
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, imageCopy, nil)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func loadImageFromBytes(data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	origin, _, err := image.Decode(reader)
	return origin, err
}

// cropCodeLocal 测试用
// benchmark 178517983 ns/op
func cropCodeLocal(imgType int, path string) {
	// 0.加载图片
	origin := loadImageFromFile(path)
	// 1. 定位小程序码位置
	w, h, dx := positionDetectionNew(imgType)
	// 2. 计算裁剪坐标
	x0, y0, l := ClaCoordinate(w, h, dx, imgType)
	// 3. 裁剪
	imageCopy, err := clip(origin, x0, y0, l, l)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "裁剪图片"}).Error(err)
	}
	// 4. 保存
	err = saveImage("D:/wlinno/qrcode_new.png", imageCopy)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "保存图片"}).Error(err)
	}
}

// ClaCoordinate 根据小程序码位置计算裁剪坐标和宽度
func ClaCoordinate(w, h, dx, imgType int) (int, int, int) {
	var (
		hr, wr, lr float64
	)
	// 根据不同类型 选择不同裁剪位置

	switch imgType {
	case ImgTypeWeChatMiniProgram:
		hr = 0.4
		wr = 0.35
		lr = 1.8
	case ImgTypeWeChat:
		hr = 0.1
		wr = 0.1
		lr = 1.4
	}

	x := h - int(float64(dx)*hr)
	y := w - int(float64(dx)*wr)
	l := int(float64(dx) * lr)
	fmt.Printf("x:%v y:%v l:%v \n", x, y, l)
	return x, y, l
}

// loadImageFromFile 从文件加载图片
func loadImageFromFile(path string) image.Image {
	file2, err := os.Open(path)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "加载图片"}).Error(err)
	}
	origin, _, err := image.Decode(file2)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "解析图片"}).Error(err)
	}
	return origin
}

// clip 图片裁剪 x y 为起点坐标 w h 为裁剪长宽
func clip(src image.Image, x, y, w, h int) (image.Image, error) {

	var subImg image.Image

	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) // 图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) // 图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) // 图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.CMYK); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.CMYK)
	} else if rgbImg, ok := src.(*image.RGBA64); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA64)
	} else if rgbImg, ok := src.(*image.Paletted); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.Paletted)
	} else {
		return src, nil
	}
	return subImg, nil
}

// saveImage 保存到文件
func saveImage(p string, src image.Image) error {

	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return err
	}
	defer f.Close()
	ext := filepath.Ext(p)

	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {

		err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})

	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(f, src)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	}
	return err
}

// isBlack 当前像素点RGB值是都为0 即黑色
func isBlack(img [][][]uint8, h, w int) bool {
	if img[h][w][0] == 0 && img[h][w][1] == 0 && img[h][w][2] == 0 {
		return true
	}
	return false
}

// positionDetection 查找二维码位置
/*
主要根据二维码三个定位点相对位置(左上 左下 右上)和大致距离才确定二维码位置
返回值1 H坐标2W坐标3偏移距离(主要用来判断二维码大小)
*/
func positionDetectionNew(imgType int) (int, int, int) {
	//如果读取出错会panic,返回图像矩阵img
	//img[height][width][4],height为图像高度,width为图像宽度
	//img[height][width][4]为第height行第width列上像素点的RGBA数值数组，值范围为0-255
	//如img[150][20][0]是150行20列处像素的红色值,img[150][20][1]是150行20列处像素的绿
	//色值，img[150][20][2]是150行20列处像素的蓝色值,img[150][20][3]是150行20列处像素
	//的alpha数值,一般用作不透明度参数,如果一个像素的alpha通道数值为0%，那它就是完全透明的.
	// 纯黑为000 纯白为255 255 255 绿色0 255 0
	var (
		minH, minW int
	)

	img := imgo.MustRead(QrCode)
	switch imgType {
	case ImgTypeWeChatMiniProgram:
		// 小程序码
		minH = int(float64(len(img)) * 0.5)
		minW = int(float64(len(img[0])) * 0.1)
	case ImgTypeWeChat:
		// 微信二维码
		minH = int(float64(len(img)) * 0.4)
		minW = int(float64(len(img[0])) * 0.5)
	default:
		logrus.WithFields(logrus.Fields{"scene": "裁剪图片-类型错误"}).Error()
		return 0, 0, 0
	}

	for h, ws := range img {
		// 二维码肯定是在一半以下
		if h < minH {
			continue
		}
		for w, _ := range ws {
			if isBlack(img, h, w) {
				for i := 1; i < len(img[h])-1; i++ {

					// 微信二维码宽度是在图片十分之一左右
					if i < minW {
						continue
					}

					nh := h + i
					nw := w + i
					if nh > len(img)-1 {
						nh = len(img) - 1
					}
					if nw > len(img[h])-1 {
						nw = len(img[h]) - 1
					}

					if (img[nh][w][0] == 0 && img[nh][w][1] == 0 && img[nh][w][2] == 0) &&
						(img[h][nw][0] == 0 && img[h][nw][1] == 0 && img[h][nw][2] == 0) {
						// 会有多个满足条件的点 随便选一个都可以
						fmt.Printf("待选点 h:%v w:%v i:%v \n", h, w, i)
						return h, w, i
					}
				}
			}
		}
	}
	return 0, 0, 0
}
