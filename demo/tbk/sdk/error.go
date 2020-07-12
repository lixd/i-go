package gotbk

import "errors"

var (
	AppKeyNil    = errors.New("appKey 不能为空")
	AppSecretNil = errors.New("appSecret 不能为空")
	RouterNil    = errors.New("router 不能为空")
)
