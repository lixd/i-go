package controller

import (
	"github.com/gin-gonic/gin"
	resultcode2 "i-go/demo/constant/resultcode"
	"i-go/demo/model"
	"i-go/demo/server"
	"log"
	"net/http"
	"strings"
)

type Controller interface {
	LoginHandler(c *gin.Context)
	LogoutHandler(c *gin.Context)
	RegisterHandler(c *gin.Context)
	RestPwdHandler(c *gin.Context)
	ForgetPwdHandler(c *gin.Context)
}
type AccountController struct {
	Server *server.Server
}

// 登录接口
func (ac *AccountController) LoginHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		log.Printf("Params Bind Err =%v ", err)
		c.JSON(http.StatusBadRequest, &model.Response{
			Code: resultcode2.Fail,
			Data: err,
			Msg:  "参数错误"})
		return
	}
	log.Printf("user.Phone=%v user.Password=%v", user.Phone, user.Password)
	user.Password = strings.TrimSpace(user.Password)
	user.Phone = strings.TrimSpace(user.Phone)
	// result, err := server.AccountServer{}.LoginServer(user.Phone, user.Password)
	result, err := (*ac.Server).LoginServer(user.Phone, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, result)
	}
	c.JSON(http.StatusOK, result)
}

func (ac *AccountController) RegisterHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{
			Code: resultcode2.Fail,
			Msg:  "参数错误~"})
	}
	result, err := (*ac.Server).RegisterServer(user.Phone, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{
			Code: resultcode2.Fail,
			Data: err,
			Msg:  "参数错误~"})
	}
	c.JSON(http.StatusOK, result)
}
