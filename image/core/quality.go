package core

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/pkg/errors"
)

// Quality 图片质量调整
func Quality(origin image.Image, q int) (image.Image, error) {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, origin, &jpeg.Options{Quality: q})
	if err != nil {
		return nil, errors.Wrap(err, "encode")
	}
	reader := bytes.NewReader(buf.Bytes())
	decode, _, err := image.Decode(reader)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}
	return decode, nil
}
