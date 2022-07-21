// Package srv Service 层返回值
package srv

import "i-go/demo/common/ret"

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Unauthorized 未授权
func Unauthorized() *Result {
	return &Result{
		Code: ret.Unauthorized,
		Msg:  ret.UnauthorizedMsg,
	}
}

// BadRequest 请求体语法错误，参数错误
func BadRequest() *Result {
	return &Result{
		Code: ret.Fail,
		Msg:  ret.BadRequest,
	}
}

// Forbidden 禁止访问 封禁
func Forbidden() *Result {
	return &Result{
		Code: ret.Forbidden,
		Msg:  ret.ForbiddenMsg,
	}
}

func Success(data interface{}, msg ...string) *Result {
	var message string
	if len(msg) == 0 {
		message = ret.SuccessMsg
	} else {
		message = msg[0]
	}
	return &Result{
		Code: ret.Success,
		Data: data,
		Msg:  message,
	}
}

func Fail(data interface{}, msg ...string) *Result {
	var message string
	if len(msg) == 0 {
		message = ret.FailMsg
	} else {
		message = msg[0]
	}
	return &Result{
		Code: ret.Fail,
		Data: data,
		Msg:  message,
	}
}
