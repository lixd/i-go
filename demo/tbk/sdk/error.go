package sdk

import "errors"

var (
	AppKeyNil    = errors.New("AppKey 不能为空")
	AppSecretNil = errors.New("AppSecret 不能为空")
	RouterNil    = errors.New("Router 不能为空")
)
