package core

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"
)

// 简单马赛克算法: 将图片指定范围坐标点的颜色值用中心点颜色值替换。

// Overall 全局马赛克
func Overall(origin image.Image, mx int, my int) image.Image {
	bounds := origin.Bounds()
	output := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for x := bounds.Min.X; x <= bounds.Max.X; x++ {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			if x%mx == 0 && y%my == 0 {
				var r, g, b, a uint32
				if x+mx/2 > bounds.Min.X {
					// 超过原图的部分就取当前位置颜色
					r, g, b, a = origin.At(x, y).RGBA() // 取该区域中间位置的颜色值以替换该区域所有位置颜色值
				} else {
					r, g, b, a = origin.At(x+mx/2, y+my/2).RGBA() // 取该区域中间位置的颜色值以替换该区域所有位置颜色值
				}
				for dx := 0; dx < mx; dx++ {
					for dy := 0; dy < my; dy++ {
						/* NRGBA类型代表没有预乘alpha通道的32位RGB色彩，Red、Green、Blue、Alpha各8位。
						*  RGBA类型代表传统的预乘了alpha通道的32位RGB色彩，Red、Green、Blue、Alpha各8位。
						*  Alpha类型代表一个8位的alpha通道
						*  RGBA 转 NRGBA 需要除去alpha通道，即右移8位
						 */
						output.SetNRGBA(x+dx, y+dy, color.NRGBA{
							R: uint8(r >> 8),
							G: uint8(g >> 8),
							B: uint8(b >> 8),
							A: uint8(a >> 8),
						})
					}
				}
			}
		}
	}
	return output
}

// Partial 局部马赛克 随机一个(mx,my)的区域进行马赛克处理(大小为固定为5*5),其他地方使用原本的颜色值
func Partial(origin image.Image, mx, my int) image.Image {
	var c = 5
	bounds := origin.Bounds()
	rand.Seed(time.Now().UnixNano())
	rx := rand.Intn(bounds.Max.X - mx)
	ry := rand.Intn(bounds.Max.Y - my)
	output := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if x >= rx && x < rx+mx && y >= ry && y < ry+my {
				// 满足c的倍数或者在边界时才进行马赛克处理
				if x%c == 0 && y%c == 0 || x == rx || y == ry {
					r, g, b, a := origin.At(x+c/2, y+c/2).RGBA() // 取该区域中间位置的颜色值以替换该区域所有位置颜色值
					for dx := 0; dx < c; dx++ {
						for dy := 0; dy < c; dy++ {
							/* NRGBA类型代表没有预乘alpha通道的32位RGB色彩，Red、Green、Blue、Alpha各8位。
							*  RGBA类型代表传统的预乘了alpha通道的32位RGB色彩，Red、Green、Blue、Alpha各8位。
							*  Alpha类型代表一个8位的alpha通道
							*  RGBA 转 NRGBA 需要除去alpha通道，即右移8位
							 */
							output.SetNRGBA(x+dx, y+dy, color.NRGBA{
								R: uint8(r >> 8),
								G: uint8(g >> 8),
								B: uint8(b >> 8),
								A: uint8(a >> 8),
							})
						}
					}
				}
			} else {
				// 其他位置就去对应位置颜色值
				r, g, b, a := origin.At(x, y).RGBA()
				output.SetNRGBA(x, y, color.NRGBA{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: uint8(a >> 8),
				})
			}

		}
	}
	return output
}

// averageColor 计算平均RGB颜色值
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	// 循环得到总共的红、绿、蓝颜色有多少
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64((bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y))
	fmt.Println(totalPixels)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}
