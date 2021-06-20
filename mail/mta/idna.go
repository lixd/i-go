package mta

import (
	"unicode"

	"golang.org/x/net/idna"
)

// IDNA Internationalizing Domain Names in Applications 应用程序中的国际化域名

// IsAllASCII 检测是否全为 ASCII 字符
func IsAllASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// ToASCII 转换为 ASCII 字符 主要用于将中文网址转换为国际网址
func ToASCII(domain string) (string, error) {
	return idna.ToASCII(domain)
}
