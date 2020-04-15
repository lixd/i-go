package main

import (
	"bytes"
	"fmt"
	"github.com/Comdex/imgo"
	lgozxing "github.com/lixd/gozxing"
	lqrcode "github.com/lixd/gozxing/qrcode"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
)

// 二维码裁剪工具包
const (
	QrCode = "D:/wlinno/qrcode_crop/qrcode_wechat7.png"
	//QrCode     = "D:/wlinno/qrcode_crop/qrcode_jd11.png"
	QrCodeSave = "D:/wlinno/qrcode_crop/qrcode_crop2.jpg"
	// 图片类型 小程序1 微信2
	ImgTypeWeChatMiniProgram = 1
	ImgTypeWeChat            = 2
	WidthLimit               = 480  //图片缩放宽度
	WidthLimitMin            = 1280 //图片缩放宽度
	Binaryzation             = 127
)

func main() {
	//cropCodeLocal(QrCode, QrCodeSave, ImgTypeWeChatMiniProgram)

	imgBytes := load2Bytes(QrCode)
	//imgBytes = iResize(imgBytes, ImgTypeWeChatMiniProgram)
	imgCropBytes := cropCode(imgBytes, ImgTypeWeChat)
	saveFile(imgCropBytes, QrCodeSave)
}

// getLocation 找到image中的定位点
func getLocation(image image.Image) []lgozxing.ResultPoint {
	reader := lqrcode.NewQRCodeReader()
	codeReader := reader.(*lqrcode.QRCodeReader)
	bitmap, _ := lgozxing.NewBinaryBitmapFromImage(image)
	location := codeReader.Location(bitmap, nil)
	return location
}

func load2Bytes(path string) []byte {
	file2, err := os.Open(path)
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
func cropCode(imgData []byte, imgType int) []byte {
	// 0.大图片缩放处理 暂时不需要缩放了
	//imgDataNew := iResize(imgData, imgType)
	//// 0.加载图片
	origin, err := loadImageFromBytes(imgData)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "解析图片"}).Error(err)
		return nil
	}
	var (
		w, h, dx int
	)
	/*
		<script type="text/javascript">var jd_union_unid="1002590618",jd_ad_ids="508:6",jd_union_pid="CP669ouXLhCao4neAxoAIOL665YLKgA=";var jd_width=0;var jd_height=0;var jd_union_euid="";var p="ABMGVB9cEQAQA2VEH0hfIlgRRgYlXVZaCCsfSlpMWGVEH0hfIgAJRRIVdUgHNxIPS0NyZiV9IEtASmdZF2sXAxMGURxeEwUaN1UaWhYGGgZSG1IlMk1DCEZrXmwTNwpfBkgyEgNcHF0QBRoOXRNfETITN2Ur";</script><script type="text/javascript" charset="utf-8" src="//u-x.jd.com/static/js/auto.js"></script>
	*/
	// 1. 定位小程序码位置
	switch imgType {
	case ImgTypeWeChatMiniProgram:
		w, h, dx = positionDetectionMiniProgram(imgData)
	case ImgTypeWeChat:
		//w, h, dx = positionDetectionQrCode(imgData)
		w, h, dx = positionDetectionQrCodeNew(origin)
	}
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

func positionDetectionQrCodeNew(image image.Image) (int, int, int) {
	location := getLocation(image)
	if len(location) < 3 {
		return 0, 0, 0
	}
	// location[0] 左下 [1]左上 [2]右上 [3]右下单独的小点
	w := int(location[1].GetY())
	h := int(location[1].GetX())
	dx := int(math.Abs(location[2].GetX() - location[1].GetX()))
	return w, h, dx
}

// positionDetectionQrCode 三点定位二维码位置
/*
找出所有合适的点 取距离最大的 兼容性提升的同时效率大幅降低 图片分辨率越高越慢
主要根据二维码三个定位点相对位置(左上 左下 右上)和大致距离才确定二维码位置
返回值1 H坐标2W坐标3偏移距离(主要用来判断二维码大小)
Benchmark
467*538 174ms
720*1280 460ms
1080*2340 2060ms
增加缩放后 200ms左右
*/
func positionDetectionQrCode(imgData []byte) (int, int, int) {

	var (
		tempW, tempH, maxI int
	)
	img := MustRead(imgData)
	img = imgo.Binaryzation(img, Binaryzation) // 二值化处理
	for h, ws := range img {
		for w, _ := range ws {
			if isBlack(img, h, w) {
				for i := 1; i < len(img[h])-1; i++ {
					nh, nw := transform(h, w, i, 1, img)
					//和小程序码区别是这里有三个条件 要判断右下角
					if isBlack(img, nh, w) && isBlack(img, h, nw) && isBlack(img, nh, nw) {
						// 会有多个满足条件的点 找距离最大的
						if i > maxI {
							//fmt.Printf("待选点 h:%v w:%v i:%v \n", h, w, i)
							tempH = h
							tempW = w
							maxI = i
						}
					}
				}
			}
		}
	}
	return tempH, tempW, maxI
}

// positionDetectionMiniProgram 小程序码定位
func positionDetectionMiniProgram(imgData []byte) (int, int, int) {

	var (
		tempW, tempH, maxI int
	)
	img := MustRead(imgData)
	//img = imgo.Binaryzation(img, 127) // 二值化处理
	for h, ws := range img {
		for w, _ := range ws {
			if isBlack(img, h, w) {
				for i := 1; i < len(img[h])-1; i++ {
					nh, nw := transform(h, w, i, 1, img)
					//和微信二维码区别是这里只有两个条件 不判断右下角
					if isBlack(img, nh, w) && isBlack(img, h, nw) {
						// 会有多个满足条件的点 找距离最大的
						if i > maxI {
							fmt.Printf("待选点 h:%v w:%v i:%v \n", h, w, i)
							tempH = h
							tempW = w
							maxI = i
						}
					}
				}
			}
		}
	}
	return tempH, tempW, maxI
}

func isRound(img [][][]uint8, w, h, dx, nw, nh int) (bool, int) {
	for i := 1; i < dx/2; i++ {
		if round(img, w, h, i) && round(img, nw, h, i) && round(img, w, nh, i) {
			return true, i
		}
	}
	return false, 0
}

func round(img [][][]uint8, w, h, dx int) bool {
	if isBlack(img, w, h) && isBlack(img, w-dx, h) && isBlack(img, w+dx, h) && isBlack(img, w, h-dx) && isBlack(img, w, h+dx) {
		return true
	}
	return false
}

func loadImageFromBytes(data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	origin, _, err := image.Decode(reader)
	return origin, err
}

// cropCodeLocal 测试用
// benchmark 178517983 ns/op
func cropCodeLocal(loadPath, savePath string, imgType int) {
	// 0.加载图片
	origin := loadImageFromFile(loadPath)
	imgData := loadFile(loadPath)
	//// 0.大图片缩放
	//imgDataResize := iResize(imgData)
	var (
		w, h, dx int
	)
	// 1. 定位小程序码位置
	switch imgType {
	case ImgTypeWeChatMiniProgram:
		w, h, dx = positionDetectionMiniProgram(imgData)
	case ImgTypeWeChat:
		w, h, dx = positionDetectionQrCode(imgData)
	}
	// 2. 计算裁剪坐标
	x0, y0, l := ClaCoordinate(w, h, dx, imgType)
	// 3. 裁剪
	imageCopy, err := clip(origin, x0, y0, l, l)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "裁剪图片"}).Error(err)
	}
	// 4. 保存
	err = saveImage(savePath, imageCopy)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "保存图片"}).Error(err)
	}
}

// 图片缩放 输入输出都为[]byte
func iResize(imgData []byte, imgType int) []byte {
	// 小程序码不缩放

	reader := bytes.NewReader(imgData)
	img, _, _ := image.Decode(reader)
	reader2 := bytes.NewReader(imgData)
	conf, _, _ := image.DecodeConfig(reader2)
	if imgType == ImgTypeWeChat {
		// 超过指定大小则缩放
		if conf.Width > WidthLimit {
			rate := float64(conf.Height) / float64(conf.Width)
			img = ImageResize(img, WidthLimit*2, int(WidthLimit*2*rate))
			//saveImage(QrCodeSave, img)
		}
	} else if imgType == ImgTypeWeChatMiniProgram {
		// 小程序码放大
		if conf.Height < WidthLimitMin {
			rate := float64(conf.Height) / float64(conf.Width)
			img = ImageResize(img, int(WidthLimitMin/rate), WidthLimitMin)
			//saveImage(QrCodeSave, img)
		}
	}

	buffer := new(bytes.Buffer)
	_ = jpeg.Encode(buffer, img, nil)
	imgData = buffer.Bytes()
	return imgData
}

// ClaCoordinate 根据小程序码位置计算裁剪坐标和宽度
func ClaCoordinate(w, h, dx, imgType int) (int, int, int) {
	var (
		hr, wr, lr float64
		transform  int
	)
	// 根据不同类型 选择不同裁剪位置

	switch imgType {
	case ImgTypeWeChatMiniProgram:
		// 小程序码还是需要按比例裁剪才行
		transform = 1
		hr = 0.3
		wr = 0.3
		lr = 1.6
		//lr = 1
	case ImgTypeWeChat:
		// 微信则不需要 周围设置一定空隙即可
		hr = 0.1
		wr = 0.1
		lr = 1.2
		transform = 5
	}

	x := h - int(float64(dx)*hr) - transform
	y := w - int(float64(dx)*wr) - transform
	l := int(float64(dx)*lr) + transform*2
	fmt.Printf("x坐标:%v y坐标:%v 边长:%v \n", x, y, l)
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

// loadImageFromFile 从文件加载图片
func loadFile(path string) []byte {
	file2, err := os.Open(path)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "加载图片"}).Error(err)
	}
	data, err := ioutil.ReadAll(file2)
	if err != nil {
		logrus.WithFields(logrus.Fields{"scene": "加载图片"}).Error(err)
	}
	return data
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
	if h < 0 {
		h = 0
	}
	if h > len(img)-1 {
		h = len(img) - 1
	}
	if w < 0 {
		w = 0
	}
	if w > len(img[0])-1 {
		w = len(img[0]) - 1
	}

	if img[h][w][0] == 0 && img[h][w][1] == 0 && img[h][w][2] == 0 {
		return true
	}
	return false
}

// isGreen 当前像素点RGB值为0、255、0 即绿色
func isGreen(img [][][]uint8, h, w int) bool {
	if img[h][w][0] == 0 && img[h][w][1] == 255 && img[h][w][2] == 0 {
		return true
	}
	return false
}

// isGreen 当前像素点RGB值是都为255 即白色
func isWhite(img [][][]uint8, h, w int) bool {
	if h < 0 {
		h = 0
	}
	if h > len(img)-1 {
		h = len(img) - 1
	}
	if w < 0 {
		w = 0
	}
	if w > len(img[0])-1 {
		w = len(img[0]) - 1
	}

	if img[h][w][0] == 255 && img[h][w][1] == 255 && img[h][w][2] == 255 {
		return true
	}
	return false
}

// transform 平移指定距离
func transform(h int, w int, dis int, pType int, img [][][]uint8) (int, int) {
	var (
		nh, nw int
	)
	switch pType {
	case 1:
		nh = h + dis
		nw = w + dis
	case 2:
		nh = h - dis
		nw = w + dis
	case 3:
		nh = h + dis
		nw = w - dis
	}
	//nh := h + dis
	//nw := w + dis
	if nh > len(img)-1 {
		nh = len(img) - 1
	}
	if nw > len(img[h])-1 {
		nw = len(img[h]) - 1
	}
	return nh, nw
}

// ImageResize 缩放 w,h 为缩放后图片宽高
func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
