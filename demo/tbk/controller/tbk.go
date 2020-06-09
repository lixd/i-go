package controller

import (
	"github.com/gin-gonic/gin"
	"i-go/demo/ret"
	"i-go/demo/tbk/dto"
	"i-go/demo/tbk/server"
	"net/http"
)

type ITBK interface {
	FindURLByKeyWords(c *gin.Context)
}

type tbk struct {
	Server server.ITBK
}

func NewTBK(ser server.ITBK) ITBK {
	return &tbk{Server: ser}
}

// FindURLByKeyWords 根据关键字、链接查询商品
func (tbk *tbk) FindURLByKeyWords(c *gin.Context) {
	var m dto.TBKReq
	if err := c.ShouldBindQuery(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := tbk.Server.FindURLByKeyWords(&m)
	c.JSON(http.StatusOK, result)
}
