package controller

import (
	"github.com/gin-gonic/gin"
	"i-go/demo/account/dto"
	"i-go/demo/account/server"
	"i-go/demo/cmodel"
	"i-go/demo/ret"
	"net/http"
)

type IAccount interface {
	Insert(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Find(c *gin.Context)
	FindList(c *gin.Context)
}

type account struct {
	Server server.IAccount
}

func NewAccount(ser server.IAccount) IAccount {
	return &account{Server: ser}
}

// Insert
func (a *account) Insert(c *gin.Context) {
	var m dto.AccountReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := a.Server.Insert(&m)
	c.JSON(http.StatusOK, result)
}

// Delete
func (a *account) Delete(c *gin.Context) {
	var m dto.AccountReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := a.Server.DeleteByUserId(&m)
	c.JSON(http.StatusOK, result)
}

// Update
func (a *account) Update(c *gin.Context) {
	var m dto.AccountReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	result := a.Server.UpdateById(&m)
	c.JSON(http.StatusOK, result)
}

// Find
func (a *account) Find(c *gin.Context) {
	var m dto.AccountReq
	if err := c.ShouldBindQuery(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	res := a.Server.FindByUserId(&m)
	c.JSON(http.StatusOK, res)
}

// Find
func (a *account) FindList(c *gin.Context) {
	var m cmodel.PageModel
	if err := c.ShouldBindQuery(&m); err != nil {
		c.JSON(http.StatusOK, ret.Fail("", "参数错误"))
		return
	}
	res := a.Server.Find(&m)
	c.JSON(http.StatusOK, res)
}
