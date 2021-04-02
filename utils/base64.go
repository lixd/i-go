package utils

import "strings"

// AddBase64Header 图片base64编码 添加Header
// completeBase64-->data:image/png;base64,iVBORw0KGgo...
func AddBase64Header(base64Body, contentType string) (completeBase64 string) {
	str := []string{"data:", contentType, ";base64,", base64Body}
	completeBase64 = strings.Join(str, "")
	return completeBase64
}

// TrimBase64Header 移除Base64头
// data:image/png;base64,iVBORw0KGgo... 去掉`,`之前的部分
func TrimBase64Header(completeBase64 string) (base64Body string) {
	index := strings.Index(completeBase64, ",")
	base64Body = completeBase64[index+1:]
	return base64Body
}
