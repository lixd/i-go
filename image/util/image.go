package util

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/png"
)

// matrix2Bytes 矩阵转成[]byte
func Matrix2Bytes(imgMatrix [][][]uint8) ([]byte, error) {
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
