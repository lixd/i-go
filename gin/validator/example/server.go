package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SignUpParam struct {
	Age        uint8  `json:"age"  binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
	// 需要使用自定义校验方法checkDate做参数校验的字段Date
	Date string `json:"date" binding:"required,datetime=2006-01-02,checkDate"`
}

func main() {
	r := gin.Default()
	err := InitTrans("zh")
	Register()
	if err != nil {
		panic(err)
	}
	r.POST("/signup", func(c *gin.Context) {
		var u SignUpParam
		if err := c.ShouldBind(&u); err != nil {
			// translate all error at once
			errs := err.(validator.ValidationErrors)
			tranErrs := errs.Translate(trans)
			// 移除 结构体名
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(tranErrs),
			})
			return
		}
		// 保存入库等业务逻辑代码...
		c.JSON(http.StatusOK, "success")
	})

	_ = r.Run(":8999")
}

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
