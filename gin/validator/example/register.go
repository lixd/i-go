package main

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 自定义错误提示信息的字段名
func Register() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(JsonTag)
		v.RegisterStructValidation(SignUpParamStructLevelValidation, &SignUpParam{})
		if err := v.RegisterValidation("checkDate", customFunc); err != nil {
			fmt.Println(err)
		}
	}
}
