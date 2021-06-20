package mta

import "net/textproto"

// IsPermanent 检测是否为永久性错误
// 主要根据错误码区分 5XX 错误则为永久性错误
// rfc 文档https://tools.ietf.org/html/rfc5321#section-4.2.1
func IsPermanent(err error) bool {
	terr, ok := err.(*textproto.Error)
	if !ok {
		return false
	}
	return terr.Code >= 500 && terr.Code < 600
}
