package utils

import (
	"path/filepath"
	"strings"
)

/*
path 和 filepath 包的区别：
path包：
	path包提供了斜线分隔路径的常规操作。但是，它只能用来处理正斜杠(/)分隔的路径。不能处理带有驱动符或反斜杠(\)的Windows路径，要操作不同操作系统的路径，请使用path/filepath包。
filepath包：
	filepath包提供了兼容操作系统的路径操作。它使用正斜杠还是反斜杠取决于操作系统。处理总是使用正斜杠的路径(如URL)，请参见path包。
结论：
	在处理路径时，应尽量使用filepath包，处理url时，使用path包。
*/

// GetFilePrefix 获取文件名(不带后缀)
func GetFilePrefix(filename string) string {
	base := filepath.Base(filename)
	suffix := filepath.Ext(filename)
	return strings.TrimSuffix(base, suffix)
}

// GetFileExt 获取文件扩展名(后缀)
func GetFileExt(filename string) string {
	return filepath.Ext(filename)
}
