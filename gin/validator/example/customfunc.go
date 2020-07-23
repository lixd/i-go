package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"time"
)

// https://github.com/go-playground/validator/issues/633#issuecomment-654382345
/*
FROM
 {
  "User.Email": "Email must be a valid email address",
  "User.FirstName": "FirstName is a required field"
}
TO
{
  "Email": "Email must be a valid email address",
  "FirstName": "FirstName is a required field"
}
*/
// removeTopStruct 移除结构体名
// from struct.field to field e.g.: from User.Email to Email
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

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

// JsonTag
func JsonTag(field reflect.StructField) string {
	return field.Tag.Get("json")
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(SignUpParam)
	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}

// customFunc 自定义字段级别校验方法
func customFunc(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}
