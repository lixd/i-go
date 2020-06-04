package controller

import (
	"github.com/gin-gonic/gin"
	"i-go/demo/cmodel"
	"i-go/demo/order/dto"
	"i-go/demo/order/server"
	"i-go/demo/ret"
	"net/http"
)

type IOrder interface {
	Insert(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Find(c *gin.Context)
	FindList(c *gin.Context)
	FindOrderAndUser(c *gin.Context)
}

type order struct {
	Server server.IOrder
}

func NewOrder(ser server.IOrder) IOrder {
	return &order{Server: ser}
}

// Insert
func (u *order) Insert(c *gin.Context) {
	var m dto.OrderReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := u.Server.Insert(&m)
	c.JSON(http.StatusOK, result)
}

// Delete
func (u *order) Delete(c *gin.Context) {
	var m dto.OrderReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := u.Server.DeleteById(&m)
	c.JSON(http.StatusOK, result)
}

// Update
func (u *order) Update(c *gin.Context) {
	var m dto.OrderReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := u.Server.UpdateById(&m)
	c.JSON(http.StatusOK, result)
}

// Find
func (u *order) Find(c *gin.Context) {
	var m dto.OrderReq
	if err := c.ShouldBindQuery(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	res := u.Server.FindById(&m)
	c.JSON(http.StatusOK, res)
}

// Find
func (u *order) FindList(c *gin.Context) {
	var m cmodel.PageModel
	if err := c.ShouldBindQuery(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	res := u.Server.Find(&m)
	c.JSON(http.StatusOK, res)
}

// Find
func (u *order) FindOrderAndUser(c *gin.Context) {
	res := u.Server.FindOrderAndUser()
	c.JSON(http.StatusOK, res)
}
