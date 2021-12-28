package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"i-go/gin/swagger/ret"
)

// Hello 测试
// @Summary      测试SayHello
// @Description  向你说Hello
// @Tags         测试
// @Accept       json
// @Produce      json
// @Param        who  query     string  true             "人名"
// @Success      200  {string}  string  "{"msg": "hello  lixd"}"
// @Failure      400  {string}  string  "{"msg": "who    are  you"}"
// @Router       /hello [get]
func Hello(c *gin.Context) {
	who := c.Query("who")

	if who == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "who are u?"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "hello " + who})
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResp struct {
	Token string `json:"token"`
}

// Login 登陆
// @Summary      登陆
// @Tags         登陆注册
// @Description  登入
// @Accept       json
// @Produce      json
// @Param        user  body      LoginReq                    true  "用户名密码"
// @Success      200   {object}  ret.Result{data=LoginResp}  "token"
// @Failure      400   {object}  ret.Result                  "错误提示"
// @Router       /login [post]
func Login(c *gin.Context) {
	var m LoginReq
	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, ret.Fail("参数错误"))
		return
	}

	if m.Username == "admin" && m.Password == "123456" {
		resp := LoginResp{Token: strconv.Itoa(int(time.Now().Unix()))}
		c.JSON(http.StatusOK, ret.Success(resp))
		return
	}
	c.JSON(http.StatusUnauthorized, ret.Fail("user  or  password  error"))
}
