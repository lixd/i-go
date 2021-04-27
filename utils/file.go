package utils

import (
	"path"
	"runtime"
	"strings"
)

// GetFilePrefix 文件名(不带后缀)
func GetFilePrefix(filename string) string {
	if runtime.GOOS == "windows" {
		filename = strings.ReplaceAll(filename, "\\", "/")
	}
	base := path.Base(filename)
	suffix := path.Ext(filename)
	return strings.TrimSuffix(base, suffix)
}

// GetFileSuffix 文件后缀
func GetFileSuffix(filename string) string {
	if runtime.GOOS == "windows" {
		filename = strings.ReplaceAll(filename, "\\", "/")
	}
	return path.Ext(filename)
}
