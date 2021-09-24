package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"i-go/core/http/ret/svc"
	"i-go/demo/account/dto"
	"i-go/demo/account/server"
	"i-go/demo/cmodel"
)

type IAccount interface {
	Insert(c *gin.Context)
	DeleteByUserId(c *gin.Context)
	Update(c *gin.Context)
	FindByUserId(c *gin.Context)
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
	var m dto.AccountInsertReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, svc.Fail("", "参数错误"))
		return
	}
	result := a.Server.Insert(&m)
	c.JSON(http.StatusOK, result)
}

// Delete
func (a *account) DeleteByUserId(c *gin.Context) {
	var m dto.AccountReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, svc.Fail("", "参数错误"))
		return
	}
	result := a.Server.DeleteByUserId(m.UserId)
	c.JSON(http.StatusOK, result)
}

// Update
func (a *account) Update(c *gin.Context) {
	var m dto.AccountReq
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusOK, svc.Fail("", "参数错误"))
		return
	}
	result := a.Server.Update(&m)
	c.JSON(http.StatusOK, result)
}

// FindByUserId
func (a *account) FindByUserId(c *gin.Context) {
	strUserId := c.Param("userId")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		c.JSON(http.StatusOK, svc.Fail("", "参数错误"))
		return
	}

	res := a.Server.FindByUserId(uint(userId))
	c.JSON(http.StatusOK, res)
}

// Find
func (a *account) FindList(c *gin.Context) {
	var m cmodel.Page
	if err := c.ShouldBindQuery(&m); err != nil {
		c.JSON(http.StatusOK, svc.Fail("", "参数错误"))
		return
	}
	res := a.Server.FindList(&m)
	c.JSON(http.StatusOK, res)
}
