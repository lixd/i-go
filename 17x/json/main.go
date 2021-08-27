package main

import (
	"log"
	"strings"
)

func main() {
	// str2:=`<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"></head><body style=\"margin: 0;background: #ccc\"><a href=\"{{链接}}\" target=\"_blank\" style=\"text-decoration: none;color: #ccc\"><div style=\"width: 100%;min-height: 900px;background: #ccc;\">{{伪原创}}<img style=\"position: absolute;left: 0;right: 0;margin: auto;top: 10%\" src=\"{{图片}}\"></div></a></body><html>`

	str := `<div class=\"right flex1\" style=\"padding: 30px 80px;\">`
	escape := terEscape(str)
	log.Println(escape)
}

func terEscape(str string) string {
	return strings.ReplaceAll(str, `\"`, `"`)
}
