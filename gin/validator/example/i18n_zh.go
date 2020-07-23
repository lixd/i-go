package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// i18n 国际化流程
//1.定义一个 全局翻译器
//2.每次返回的时候用全局翻译器进行翻译。。

var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() // 中文翻译器

		// 第一个参数是备用（fallback）的语言环境 后面的参数是应该支持的语言环境（支持多个）
		uni := ut.New(zhT, zhT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		// 这里可以用 gin 框架自带的 validate 或者 validator 库 都是一样的
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
		/*		validate := validator.New()
				err = zhTranslations.RegisterDefaultTranslations(validate, trans)*/
		if err != nil {
			return err
		}
		return
	}
	return
}
