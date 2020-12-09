package image

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

// crowdFunding 众筹中
func Hello() {
	// 自定义字体
	// ttf, err := ioutil.ReadFile("./msyhl.ttc")
	// if err != nil {
	// 	return nil, nil, err
	// }
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return
	}
	dc := gg.NewContext(400, 230)
	face := truetype.NewFace(font, &truetype.Options{Size: 20})
	dc.SetFontFace(face)
	dc.SetHexColor("#EA5949")
	dc.DrawString("Hello world", 40, 40)
	err = dc.SavePNG("out.png")
	if err != nil {
		return
	}
	return
}
