package controller

import (
	"github.com/gin-gonic/gin"
	"i-go/gin/demo/constant/resultcode"
	"i-go/gin/demo/model"
	"i-go/gin/demo/server"
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

func (ac *AccountController) LoginHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{
			Code: resultcode.Fail,
			Msg:  "参数错误"})
		return
	}
	user.Password = strings.TrimSpace(user.Password)
	user.Phone = strings.TrimSpace(user.Phone)
	result, err := (*ac.Server).LoginServer(user.Phone, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, result)
	}
	c.JSON(http.StatusOK, result)

}
