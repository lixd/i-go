package ret

import (
	"i-go/demo/constant/retcode"
	"i-go/demo/constant/retmsg"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Unauthorized() *Result {
	return &Result{
		Code: retcode.Unauthorized,
		Msg:  retmsg.Unauthorized,
	}
}

func Success(data interface{}, msg ...string) *Result {
	var message string
	if len(msg) == 0 {
		message = retmsg.Success
	} else {
		message = msg[0]
	}
	return &Result{
		Code: retcode.Success,
		Data: data,
		Msg:  message,
	}
}

func Fail(data interface{}, msg ...string) *Result {
	var message string
	if len(msg) == 0 {
		message = retmsg.Fail
	} else {
		message = msg[0]
	}
	return &Result{
		Code: retcode.Fail,
		Data: data,
		Msg:  message,
	}
}
