package controller

import (
	"net/http"
	"strconv"

	"i-go/demo/common/ret/srv"
	"i-go/demo/order/dto"
	"i-go/demo/order/server"

	"github.com/gin-gonic/gin"
)

type IOrder interface {
	Insert(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	Find(c *gin.Context)
	FindOrderAndUser(c *gin.Context)
}

type order struct {
	Server server.IOrder
}

func NewOrder(ser server.IOrder) IOrder {
	return &order{Server: ser}
}

func (o *order) Insert(c *gin.Context) {
	var m dto.OrderReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, srv.Fail("", "参数错误"))
		return
	}
	result := o.Server.Insert(&m)
	c.JSON(http.StatusOK, result)
}

func (o *order) Delete(c *gin.Context) {
	var m dto.OrderReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, srv.Fail("", "参数错误"))
		return
	}
	result := o.Server.Delete(&m)
	c.JSON(http.StatusOK, result)
}

func (o *order) Update(c *gin.Context) {
	var m dto.OrderReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, srv.Fail("", "参数错误"))
		return
	}
	result := o.Server.Update(&m)
	c.JSON(http.StatusOK, result)
}

func (o *order) FindById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusOK, srv.Fail("", "参数错误"))
		return
	}
	res := o.Server.FindById(uint(id))
	c.JSON(http.StatusOK, res)
}

func (o *order) Find(c *gin.Context) {
	var m dto.OrderReq
	if err := c.ShouldBindQuery(&m); err != nil {
		c.JSON(http.StatusOK, srv.Fail("", "参数错误"))
		return
	}
	res := o.Server.Find(&m)
	c.JSON(http.StatusOK, res)
}

func (o *order) FindOrderAndUser(c *gin.Context) {
	res := o.Server.FindOrderAndUser()
	c.JSON(http.StatusOK, res)
}
